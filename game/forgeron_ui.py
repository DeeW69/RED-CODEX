import tkinter as tk
from tkinter import messagebox
import json, sys, os

FILENAME = "data_players.json"

# ------------------ util JSON ------------------
def load_player_data():
    if not os.path.exists(FILENAME):
        # squelette minimal si le fichier n'existe pas
        return {
            "player": {
                "stats": {
                    "name": "Héros", "level": 1, "experience": 0,
                    "gold": 0, "health": 100, "max_health": 100,
                    "mana": 50, "max_mana": 50, "spells": []
                },
                "equipment": {
                    "head": {"name":"", "level_star":1},
                    "weapon":{"name":"", "level_star":1},
                    "armor":{"name":"", "level_star":1}
                },
                "inventory": {"potions": [], "weapons": [], "drops": {}, "gold": 0},
                "companions": {}
            }
        }
    with open(FILENAME, "r", encoding="utf-8") as f:
        return json.load(f)

def save_player_data(data):
    with open(FILENAME, "w", encoding="utf-8") as f:
        json.dump(data, f, indent=2, ensure_ascii=False)

# ------------------ recettes ------------------
# clef = tuple triée des composants posés dans la table
RECIPES = {
    tuple(sorted(["peau", "crocs"])): {"item": "bottes en cuir", "qty": 1, "category": "drops"},
    tuple(sorted(["viande", "venin"])): {"item": "potion de poison", "qty": 1, "category": "drops"},
    tuple(sorted(["carapace", "peau", "crocs"])): {"item": "armure légère", "qty": 1, "category": "drops"},
    # ajoute tes recettes ici
}

# ------------------ état ------------------
data = load_player_data()
stats = data["player"]["stats"]
inv = data["player"]["inventory"]
drops = inv.get("drops", {})
inv["drops"] = drops  # s'assure que la clé existe

# table de craft 3x3 (chaque case = "" ou nom d'item)
craft_grid = [["" for _ in range(3)] for _ in range(3)]

selected_item = {"name": None}   # item “tenu en main” (clic inventaire → clic case craft)
selected_widget = None           # pour surligner le bouton sélectionné

# ------------------ UI ------------------
root = tk.Tk()
root.title("Forgeron & Crafting")

# grilles responsives
for col in range(4):
    root.grid_columnconfigure(col, weight=1)

# ---- Stats (gauche) ----
frame_stats = tk.LabelFrame(root, text="Stats joueur", padx=10, pady=10, bg="#d9f0ff")
frame_stats.grid(row=0, column=0, rowspan=3, sticky="nsew", padx=6, pady=6)

lbl_name   = tk.Label(frame_stats, text=f"Nom : {stats.get('name','')}", bg="#d9f0ff")
lbl_level  = tk.Label(frame_stats, text=f"Niveau : {stats.get('level',0)}", bg="#d9f0ff")
lbl_hp     = tk.Label(frame_stats, text=f"PV : {stats.get('health',0)}/{stats.get('max_health',0)}", bg="#d9f0ff")
lbl_mp     = tk.Label(frame_stats, text=f"Mana : {stats.get('mana',0)}/{stats.get('max_mana',0)}", bg="#d9f0ff")
lbl_gold   = tk.Label(frame_stats, text=f"Or : {stats.get('gold',0)}", bg="#d9f0ff")

lbl_name.grid(row=0, column=0, sticky="w")
lbl_level.grid(row=1, column=0, sticky="w")
lbl_hp.grid(row=2, column=0, sticky="w")
lbl_mp.grid(row=3, column=0, sticky="w")
lbl_gold.grid(row=4, column=0, sticky="w")

# ---- Inventaire (centre) ----
frame_items = tk.LabelFrame(root, text="Inventaire (drops)", padx=10, pady=10, bg="#fff3d1")
frame_items.grid(row=0, column=1, rowspan=3, sticky="nsew", padx=6, pady=6)

inv_buttons = {}
def refresh_inventory():
    # nettoie
    for w in frame_items.grid_slaves():
        w.destroy()
    tk.Label(frame_items, text="Cliquer pour sélectionner (1 unité)", bg="#fff3d1").grid(row=0, column=0, sticky="w")
    # recrée boutons
    r = 1
    for item, qty in sorted(drops.items()):
        txt = f"{item}  x{qty}"
        state = tk.NORMAL if qty > 0 else tk.DISABLED
        btn = tk.Button(frame_items, text=txt, state=state,
                        command=lambda it=item: on_pick_item(it),
                        relief="raised", width=22, bg="lemon chiffon")
        btn.grid(row=r, column=0, sticky="ew", pady=2)
        inv_buttons[item] = btn
        r += 1

def on_pick_item(item):
    global selected_widget
    # annule l’ancien highlight
    if selected_widget is not None:
        try:
            selected_widget.config(bg="lemon chiffon")
        except Exception:
            pass
    # si plus de stock → rien
    if drops.get(item, 0) <= 0:
        return
    # sélectionne
    selected_item["name"] = item
    selected_widget = inv_buttons.get(item)
    if selected_widget:
        selected_widget.config(bg="#ffd27d")  # highlight

# ---- Craft (droite haut) ----
frame_craft = tk.LabelFrame(root, text="Table de craft (3x3)", padx=8, pady=8, bg="#ffd1ec")
frame_craft.grid(row=0, column=2, sticky="nsew", padx=6, pady=6, columnspan=2)

craft_buttons = [[None]*3 for _ in range(3)]

def refresh_craft():
    for r in range(3):
        for c in range(3):
            txt = craft_grid[r][c] if craft_grid[r][c] else "—"
            craft_buttons[r][c].config(text=txt)

def on_craft_click(r, c):
    it = selected_item["name"]
    cur = craft_grid[r][c]
    # si un item est sélectionné et la case est libre → poser l’item (consomme 1)
    if it and not cur:
        if drops.get(it,0) > 0:
            drops[it] -= 1
            craft_grid[r][c] = it
            if drops[it] == 0:
                # désélectionne si plus de stock
                deselect()
            refresh_inventory()
            refresh_craft()
        return
    # si aucun item sélectionné et la case est occupée → récupérer l’item
    if not it and cur:
        craft_grid[r][c] = ""
        drops[cur] = drops.get(cur,0) + 1
        refresh_inventory()
        refresh_craft()
        return
    # sinon (case occupée et item sélectionné) : swap interdit → bip
    root.bell()

for r in range(3):
    for c in range(3):
        b = tk.Button(frame_craft, text="—", width=12, height=2, bg="#ffc2e8",
                      command=lambda rr=r, cc=c: on_craft_click(rr, cc))
        b.grid(row=r, column=c, padx=3, pady=3)
        craft_buttons[r][c] = b

# ---- Résultat (droite bas) ----
frame_result = tk.LabelFrame(root, text="Résultat", padx=10, pady=10, bg="#e9e6ff")
frame_result.grid(row=1, column=2, sticky="nsew", padx=6, pady=6, columnspan=2)
label_result = tk.Label(frame_result, text="?", bg="#e9e6ff", font=("Arial", 13))
label_result.pack(anchor="w")

# ---- Actions (droite bas) ----
frame_actions = tk.Frame(root, padx=10, pady=10)
frame_actions.grid(row=2, column=2, sticky="nsew", padx=6, pady=6, columnspan=2)

def deselect():
    global selected_widget
    selected_item["name"] = None
    if selected_widget is not None:
        try:
            selected_widget.config(bg="lemon chiffon")
        except Exception:
            pass
    selected_widget = None

def clear_craft(return_to_inventory=True):
    # vide la table. si return_to_inventory, rend les composants
    for r in range(3):
        for c in range(3):
            if craft_grid[r][c]:
                if return_to_inventory:
                    drops[craft_grid[r][c]] = drops.get(craft_grid[r][c],0) + 1
                craft_grid[r][c] = ""
    refresh_inventory()
    refresh_craft()
    deselect()
    label_result.config(text="?")

def validate_craft():
    # construit la clé triée des composants
    used = []
    for r in range(3):
        for c in range(3):
            if craft_grid[r][c]:
                used.append(craft_grid[r][c])
    if not used:
        label_result.config(text="❌ Aucun composant")
        return

    key = tuple(sorted(used))
    recipe = RECIPES.get(key)
    if recipe is None:
        # échec → rend les composants
        clear_craft(return_to_inventory=True)
        label_result.config(text="❌ Échec du craft (combinaison inconnue)")
        return

    # succès → composants déjà consommés au placement, on ajoute le résultat
    res_item = recipe["item"]
    qty = recipe.get("qty", 1)
    category = recipe.get("category", "drops")

    if category == "drops":
        drops[res_item] = drops.get(res_item, 0) + qty
    else:
        # extension possible: ajouter dans weapons/potions etc.
        drops[res_item] = drops.get(res_item, 0) + qty

    save_player_data(data)
    label_result.config(text=f"✅ {res_item} x{qty}")
    clear_craft(return_to_inventory=False)  # on garde la conso
    refresh_inventory()

btn_validate = tk.Button(frame_actions, text="Valider le craft", command=validate_craft, bg="#b8ffb8")
btn_clear    = tk.Button(frame_actions, text="Vider la table", command=lambda: clear_craft(True))
btn_validate.grid(row=0, column=0, padx=4)
btn_clear.grid(row=0, column=1, padx=4)

# ---- Achat spellbook ----
def buy_spellbook():
    cost = 50
    gold = stats.get("gold", 0)
    if gold < cost:
        messagebox.showwarning("Échec", "Pas assez d'or !")
        return
    stats["gold"] = gold - cost
    if "spells" not in stats:
        stats["spells"] = []
    if "Boule de feu" not in stats["spells"]:
        stats["spells"].append("Boule de feu")
    save_player_data(data)
    lbl_gold.config(text=f"Or : {stats['gold']}")
    messagebox.showinfo("Achat réussi", "📘 Vous avez acheté : Boule de feu")

btn_spell = tk.Button(root, text="Acheter SpellBook (Boule de feu) - 50 or", command=buy_spellbook)
btn_spell.grid(row=3, column=0, columnspan=3, pady=8, sticky="ew", padx=6)

# ---- Quitter : renvoie un JSON pour Go si besoin ----
def on_close():
    # renvoyer un petit état pour battle.go (facultatif)
    print(json.dumps({
        "action": "forgeron_ui",
        "remaining_gold": stats.get("gold", 0),
        "spells": stats.get("spells", []),
    }, ensure_ascii=False))
    sys.stdout.flush()
    root.destroy()

root.protocol("WM_DELETE_WINDOW", on_close)

# init affichage
refresh_inventory()
refresh_craft()

root.mainloop()
