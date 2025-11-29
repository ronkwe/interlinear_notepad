# Glossa - Interlinear Notepad

Glossa is a specialized tool designed for linguists to create, edit, and export interlinear glossed text. It provides a clean, distraction-free interface for analyzing linguistic data, with features tailored for morphological analysis and translation.

## Features

-   **Dynamic Interlinear Blocks**: Create gloss blocks with customizable lines (e.g., Source, Morphemes, Gloss, Translation).
-   **Flexible Line Management**: Add, remove, hide, or show lines globally across the document.
-   **Smart Editing**:
    -   **Space** to add a new column.
    -   **Backspace** at the start of a cell to merge with the previous column.
-   **Export Options**:
    -   **PDF**: Generate high-quality PDFs ready for sharing.
    -   **Markdown**: Export to Markdown for integration with other documents.
    -   **Native File Format**: Save and load your work using `.glossa` files.
-   **Cross-Platform**: Available as a standalone application for Linux and Windows.

## Getting Started

### Installation

Download the latest release for your operating system:
-   **Linux**: `glossa`
-   **Windows**: `glossa.exe`

### Usage

1.  **Add a Block**: Click the "+ Add Block" button to start a new sentence or phrase.
2.  **Edit Text**: Type in the cells. Use `Tab` to move between cells.
3.  **Add Columns**: Press `Space` at the end of a cell to create a new column for the next word/morpheme.
4.  **Manage Lines**: Use the toolbar to toggle line labels or add/remove lines as needed.
5.  **Export**: Use the "Export PDF" or "Export Markdown" buttons to save your work.

## Technology Stack

Glossa is built using **Wails**, combining the power of **Go** for the backend with the flexibility of **Svelte** for the frontend.

-   **Backend**: Go (File I/O, PDF Generation)
-   **Frontend**: Svelte (UI, State Management)
-   **Build Tool**: Wails v2

## License

[MIT License](LICENSE)
