<script>
    import { createEventDispatcher } from "svelte";
    import { documentStore } from "../stores.js";
    import {
        OpenFile,
        SaveFile,
        SaveMarkdownFile,
        OpenMarkdownFile,
        SavePDFFile,
    } from "../../wailsjs/go/main/App";

    const dispatch = createEventDispatcher();

    let activeMenu = null;

    function toggleMenu(menu) {
        if (activeMenu === menu) {
            activeMenu = null;
        } else {
            activeMenu = menu;
        }
    }

    function closeMenu() {
        activeMenu = null;
    }

    // File Operations
    async function handleNew() {
        $documentStore = {
            version: "1.0",
            config: { ...$documentStore.config }, // Keep config
            blocks: [
                {
                    columns: [
                        {
                            id: crypto.randomUUID(),
                            lines: Array($documentStore.config.line_count).fill(
                                "",
                            ),
                        },
                    ],
                },
            ],
        };
        closeMenu();
    }

    async function handleOpen() {
        try {
            const result = await OpenFile();
            if (result) $documentStore = result;
        } catch (err) {
            alert("Error opening file: " + err);
        }
        closeMenu();
    }

    async function handleSave() {
        try {
            const path = await SaveFile($documentStore);
            if (path) alert("Saved to: " + path);
        } catch (err) {
            alert("Error saving: " + err);
        }
        closeMenu();
    }

    async function handleImportMarkdown() {
        try {
            const result = await OpenMarkdownFile();
            if (result) $documentStore = result;
        } catch (err) {
            alert("Error importing Markdown: " + err);
        }
        closeMenu();
    }

    async function handleExportMarkdown() {
        try {
            const path = await SaveMarkdownFile($documentStore);
            if (path) alert("Exported to: " + path);
        } catch (err) {
            alert("Error exporting: " + err);
        }
        closeMenu();
    }

    async function handleExportPDF() {
        try {
            const path = await SavePDFFile($documentStore);
            if (path) alert("Exported to: " + path);
        } catch (err) {
            alert("Error exporting: " + err);
        }
        closeMenu();
    }

    // View Operations
    function toggleLabels() {
        $documentStore.config.show_labels = !$documentStore.config.show_labels;
        closeMenu();
    }

    // Window click to close menu
    function handleWindowClick(event) {
        if (!event.target.closest(".menu-bar")) {
            closeMenu();
        }
    }
</script>

<svelte:window on:click={handleWindowClick} />

<div class="menu-bar">
    <div class="menu-item">
        <button class="menu-btn" on:click={() => toggleMenu("file")}
            >File</button
        >
        {#if activeMenu === "file"}
            <div class="dropdown">
                <button on:click={handleNew}>New</button>
                <div class="separator"></div>
                <button on:click={handleOpen}>Open...</button>
                <button on:click={handleSave}>Save...</button>
                <div class="separator"></div>
                <button on:click={handleImportMarkdown}
                    >Import Markdown...</button
                >
                <button on:click={handleExportMarkdown}
                    >Export Markdown...</button
                >
                <button on:click={handleExportPDF}>Export PDF...</button>
            </div>
        {/if}
    </div>

    <div class="menu-item">
        <button class="menu-btn" on:click={() => toggleMenu("view")}
            >View</button
        >
        {#if activeMenu === "view"}
            <div class="dropdown">
                <button on:click={toggleLabels}>
                    {$documentStore.config.show_labels
                        ? "Hide Labels"
                        : "Show Labels"}
                </button>
            </div>
        {/if}
    </div>

    <div class="menu-item">
        <button class="menu-btn" on:click={() => toggleMenu("help")}
            >Help</button
        >
        {#if activeMenu === "help"}
            <div class="dropdown">
                <button
                    on:click={() => alert("Glossa v1.0\nInterlinear Notepad")}
                    >About</button
                >
            </div>
        {/if}
    </div>
</div>

<style>
    .menu-bar {
        display: flex;
        background-color: var(--bg-panel);
        border-bottom: 1px solid var(--border);
        padding: 0 10px;
        height: 35px;
        align-items: center;
        user-select: none;
    }

    .menu-item {
        position: relative;
    }

    .menu-btn {
        background: transparent;
        border: none;
        color: var(--text-main);
        padding: 5px 10px;
        cursor: pointer;
        font-size: 0.9em;
        border-radius: 4px;
    }

    .menu-btn:hover {
        background-color: var(--bg-input);
    }

    .dropdown {
        position: absolute;
        top: 100%;
        left: 0;
        background-color: var(--bg-panel);
        border: 1px solid var(--border);
        border-radius: 4px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
        min-width: 150px;
        z-index: 1000;
        padding: 5px 0;
    }

    .dropdown button {
        display: block;
        width: 100%;
        text-align: left;
        background: transparent;
        border: none;
        color: var(--text-main);
        padding: 8px 15px;
        cursor: pointer;
        font-size: 0.9em;
    }

    .dropdown button:hover {
        background-color: var(--primary);
        color: var(--bg-app);
    }

    .separator {
        height: 1px;
        background-color: var(--border);
        margin: 5px 0;
    }
</style>
