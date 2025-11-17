<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
    <img src="extension/public/icon/128.png" alt="Logo" width="80" height="80" />

<h3 align="center">LeetCode Rich Presence</h3>

  <p align="center">Discord Rich Presence for LeetCode using a browser extension and a local Go backend.</p>
</div>

 <br />

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">🧭 About The Project</a>
      <ul>
        <li><a href="#built-with">🏗️ Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">📋 Getting Started</a>
      <ul>
        <li><a href="#prerequisites">🗺️ Prerequisites</a></li>
        <li><a href="#installation">⚙️ Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">💾 Usage</a></li>
    <li><a href="#configuration">🧰 Configuration</a></li>
    <li><a href="#troubleshooting">🧯 Troubleshooting</a></li>
    <li><a href="#contributing">🔗 Contributing</a></li>
    <li><a href="#license">📰 License</a></li>
    <li><a href="#contact">📫 Contact</a></li>
    <li><a href="#acknowledgments">⛱️ Acknowledgments</a></li>
  </ol>
</details>

<br>

<!-- ABOUT THE PROJECT -->

## 🧭 About The Project

LeetCode Rich Presence updates your Discord status with the problem you’re viewing/solving on LeetCode. It consists of:
- A browser extension that detects the current LeetCode problem and difficulty.
- A local WebSocket server written in Go that forwards updates to Discord via IPC.
- A lightweight OAuth2 flow to authorize Discord RPC.

High-level flow:
- Extension observes tabs on `https://leetcode.com/problems/...`
- Sends `{ title, url }` over WebSocket to `ws://localhost:{port}`
- Backend updates Discord activity with the current problem and a running timestamp

### 🏗️ Built With

- Go (backend WebSocket server + Discord IPC)
- vite-plugin-web-extension
- webextension-polyfill
- Gorilla WebSocket

<p align="right"><a href="#top">⬆️</a></p>

<!-- GETTING STARTED -->

## 📋 Getting Started

Follow these steps to run the backend and load the extension.

Note: you could use the latest version from the release tab.

### 🗺️ Prerequisites

- Discord Desktop app installed and running
- Go (recommended latest stable)
- Node.js 18+ and your preferred package manager (pnpm/npm)
- Chromium-based browser (Chrome/Edge) with developer mode enabled
- A Discord application (to obtain `CLIENTID` and `CLIENTSECRET`) with the `rpc.activities.write` scope

### ⚙️ Installation

1) Backend (Go)

- Create a `.env` file in `backend/`:
```bash
CLIENTID=your_discord_client_id
CLIENTSECRET=your_discord_client_secret
PORT=8085
```

- Start the server:
  cd backend
  go run .  The server listens on `localhost:{PORT}` and will prompt an authorization flow the first time it connects to Discord.

2) Extension (Browser)

- Build the extension:
  cd extension
  pnpm install
  pnpm build- Load the built extension:
  - Chrome: open `chrome://extensions`, enable Developer Mode, click “Load unpacked,” select `extension/dist/`.

<p align="right"><a href="#top">⬆️</a></p>

<!-- USAGE EXAMPLES -->

## 💾 Usage

- Open the extension popup.
- Enter the backend port (default `8085`) and click “Connect.”
- Visit a LeetCode problem page (e.g., `https://leetcode.com/problems/...`).
- Your Discord status should update to “Problem Solving” with the problem title and difficulty.

Tip: The popup footer shows connection status (idle, connected, disconnected).

<p align="right"><a href="#top">⬆️</a></p>

## 🧰 Configuration

- Environment variables (in `backend/.env`):
  - `CLIENTID` (required): Your Discord app Client ID
  - `CLIENTSECRET` (required): Your Discord app Client Secret
  - `PORT` (optional, default `8085`)
- Token storage:
  - OAuth tokens are stored locally at:
    - macOS: `~/Library/Application Support/leetcode-rich-presence/discord_tokens.json`
    - Linux: `~/.config/leetcode-rich-presence/discord_tokens.json`

<p align="right"><a href="#top">⬆️</a></p>

## 🧯 Troubleshooting

- “Invalid port” in popup:
  - Ensure the backend is running and the port matches.
- No Discord update:
  - Make sure you authorize to display activity in discord settings.
  - Make sure Discord Desktop is running (RPC works only with the desktop app).
  - The first run will open an authorization prompt inside Discord; approve it.
  - Verify tokens exist at the path above; delete them to re-trigger auth if needed.
- Connection blocked:
  - The backend only accepts WebSocket connections from `chrome-extension://` origins and `localhost`.

Platform note: Current IPC implementation targets Unix domain sockets (`/tmp/discord-ipc-0` or `$TMPDIR/discord-ipc-0`), which works on macOS/Linux.

<p align="right"><a href="#top">⬆️</a></p>

<!-- CONTRIBUTING -->

## 🔗 Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right"><a href="#top">⬆️</a></p>

<!-- CONTACT -->

## 📫 Contact

Reach me at : gauron.dorian.pro@gmail.com.

Project Link: [https://github.com/Michelprogram/leetcode-rich-presence](https://github.com/Michelprogram/leetcode-rich-presence)

<p align="right"><a href="#top">⬆️</a></p>

<!-- ACKNOWLEDGMENTS -->

## ⛱️ Acknowledgments

- Discord IPC docs
- Gorilla WebSocket
- vite-plugin-web-extension
- webextension-polyfill
- Tailwind CSS
- LeetCode