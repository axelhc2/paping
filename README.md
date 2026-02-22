# 🚀 paping (Multi-OS)

<p align="center">
  <strong>TCP Port Connectivity Tester</strong><br>
  <em>By AXEL CHETAIL - <a href="https://infrawire.fr">infrawire.fr</a></em>
</p>

---

## 🌍 Languages / Langues

## <img src="https://github.com/lipis/flag-icons/raw/main/flags/4x3/fr.svg" width="20"> Français

## <img src="https://github.com/lipis/flag-icons/raw/main/flags/4x3/us.svg" width="20"> English

---

# <img src="https://github.com/lipis/flag-icons/raw/main/flags/4x3/fr.svg" width="24"> Documentation (FR)

## 💻 Compatibilité & Installation

### 🐧 Linux (Ubuntu, Debian, CentOS, etc.)

La méthode la plus rapide pour installer la version Linux :

```bash
curl -sL https://raw.githubusercontent.com/axelhc2/paping/main/install.sh | sudo bash
```

---

### 🍎 macOS (Apple Silicon & Intel)

Ouvrez votre terminal et copiez les commandes suivantes selon votre processeur.

#### Apple Silicon (Puces M1, M2, M3, M4)

```bash
sudo curl -L https://raw.githubusercontent.com/axelhc2/paping/main/bin/paping-mac-m1 -o /usr/local/bin/paping
sudo chmod +x /usr/local/bin/paping
```

#### Intel

```bash
sudo curl -L https://raw.githubusercontent.com/axelhc2/paping/main/bin/paping-mac-intel -o /usr/local/bin/paping
sudo chmod +x /usr/local/bin/paping
```

---

### 🪟 Windows

Pour Windows, téléchargez simplement l'exécutable.

**Téléchargement :** `paping.exe`
**Utilisation :** Ouvrez un terminal (CMD ou PowerShell) dans le dossier du fichier.

```powershell
.\\paping.exe 1.1.1.1 -p 80
```

---

## 💡 Pourquoi c'est structuré comme ça ?

* **Linux :** Le `curl | bash` est le standard de l'industrie pour une installation rapide.
* **macOS :** Installation dans `/usr/local/bin` pour que la commande `paping` soit reconnue globalement immédiatement.
* **Multi-langues :** Sections FR et EN incluses pour une meilleure compatibilité GitHub.

---

# <img src="https://github.com/lipis/flag-icons/raw/main/flags/4x3/us.svg" width="24"> Documentation (EN)

## 💻 Compatibility & Installation

### 🐧 Linux (Ubuntu, Debian, CentOS, etc.)

The fastest way to install the Linux version:

```bash
curl -sL https://raw.githubusercontent.com/axelhc2/paping/main/install.sh | sudo bash
```

---

### 🍎 macOS (Apple Silicon & Intel)

Open your terminal and copy the commands according to your CPU.

#### Apple Silicon (M1, M2, M3, M4 Chips)

```bash
sudo curl -L https://raw.githubusercontent.com/axelhc2/paping/main/bin/paping-mac-m1 -o /usr/local/bin/paping
sudo chmod +x /usr/local/bin/paping
```

#### Intel

```bash
sudo curl -L https://raw.githubusercontent.com/axelhc2/paping/main/bin/paping-mac-intel -o /usr/local/bin/paping
sudo chmod +x /usr/local/bin/paping
```

---

### 🪟 Windows

For Windows, simply download the executable.

**Download:** `paping.exe`
**Usage:** Open a terminal (CMD or PowerShell) in the file directory.

```powershell
.\\paping.exe 1.1.1.1 -p 80
```

---

## 👨‍💻 Credits

Developed with ❤️ by AXEL CHETAIL ([https://infrawire.fr](https://infrawire.fr))

---

## 🛠️ Documentation Improvements

1. Removed buggy `<img>` tags and replaced them with lightweight icons.
2. Clean code blocks for easy copy-paste on GitHub.
3. Clear hierarchy for automatic GitHub table of contents.
4. Full bilingual support (FR / EN) with flag sections for professional README display.
