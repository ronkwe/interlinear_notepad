<script>
  import {
    OpenFile,
    SaveFile,
    ExportToMarkdown,
    ExportToPDF,
    SaveMarkdownFile,
    SavePDFFile,
  } from "../../wailsjs/go/main/App";
  import { documentStore } from "../stores.js";

  async function openFile() {
    try {
      const result = await OpenFile();
      if (result) {
        $documentStore = result;
      }
    } catch (err) {
      console.error(err);
      alert("Error opening file: " + err);
    }
  }

  async function saveFile() {
    try {
      const path = await SaveFile($documentStore);
      if (path) {
        alert("File saved successfully to: " + path);
      }
    } catch (err) {
      console.error(err);
      alert("Error saving file: " + err);
    }
  }

  async function exportMarkdown() {
    try {
      const path = await SaveMarkdownFile($documentStore);
      if (path) {
        alert("Markdown saved successfully to: " + path);
      }
    } catch (err) {
      console.error(err);
      alert("Error exporting to Markdown: " + err);
    }
  }

  async function exportPDF() {
    try {
      const path = await SavePDFFile($documentStore);
      if (path) {
        alert("PDF saved successfully to: " + path);
      }
    } catch (err) {
      console.error(err);
      alert("Error exporting to PDF: " + err);
    }
  }

  function toggleLabels() {
    $documentStore.config.show_labels = !$documentStore.config.show_labels;
  }

  function addLine() {
    $documentStore.config.line_count++;
    $documentStore.config.line_options.push({
      label: "New Line",
      visible: true,
    });
    $documentStore.blocks.forEach((block) => {
      block.columns.forEach((col) => {
        col.lines.push("");
      });
    });
    $documentStore = $documentStore; // Trigger update
  }
</script>

<div class="toolbar">
  <button on:click={openFile}>Open</button>
  <button on:click={saveFile}>Save</button>
  <button on:click={exportMarkdown}>Export Markdown</button>
  <button on:click={exportPDF}>Export PDF</button>
  <div class="toolbar-group">
    <button on:click={toggleLabels}>
      {$documentStore.config.show_labels ? "Hide Labels" : "Show Labels"}
    </button>
    <button on:click={addLine}>+ Line</button>
  </div>
</div>

<style>
  .toolbar {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }
  button {
    padding: 8px 12px;
    border-radius: 6px;
    border: 1px solid var(--border);
    font-weight: 500;
    font-family: inherit;
    background-color: var(--bg-panel);
    color: var(--text-main);
    cursor: pointer;
    transition: all 0.25s;
  }
  button:hover {
    border-color: var(--primary);
    color: var(--primary);
    background-color: var(--bg-input);
  }
  button:focus,
  button:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }
</style>
