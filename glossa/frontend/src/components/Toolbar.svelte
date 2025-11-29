<script>
  import { documentStore } from "../stores.js";

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
  <div class="toolbar-group">
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
