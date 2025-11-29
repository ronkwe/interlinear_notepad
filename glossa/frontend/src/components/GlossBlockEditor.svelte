<script>
  import { createEventDispatcher } from "svelte";
  import { documentStore } from "../stores.js";
  export let block;

  const dispatch = createEventDispatcher();

  function handleKeydown(event, columnIndex) {
    const el = event.target;
    const selection = window.getSelection();

    // Add a new column when space is pressed at the end of a cell
    if (event.key === " " && selection.anchorOffset === el.innerText.length) {
      event.preventDefault();
      insertColumnAfter(columnIndex);
    }

    // Merge with previous column if backspace is pressed at the start of a cell
    if (
      event.key === "Backspace" &&
      selection.anchorOffset === 0 &&
      columnIndex > 0
    ) {
      event.preventDefault();
      mergeWithPreviousColumn(columnIndex);
    }
  }

  function addColumn(index) {
    const newColumn = {
      lines: Array($documentStore.config.line_count).fill(""),
    };
    block.columns.splice(index + 1, 0, newColumn);
    block.columns = block.columns; // Trigger Svelte reactivity
    // TODO: Focus on the first line of the new column
  }

  function mergeWithPreviousColumn(index) {
    const currentColumn = block.columns[index];
    const previousColumn = block.columns[index - 1];

    currentColumn.lines.forEach((line, i) => {
      previousColumn.lines[i] += line;
    });

    block.columns.splice(index, 1);
    block.columns = block.columns; // Trigger Svelte reactivity
    // TODO: Set cursor position correctly in the merged cell
  }

  function deleteBlock() {
    dispatch("delete");
  }

  function deleteLine(index) {
    if ($documentStore.config.line_count > 1) {
      $documentStore.config.line_count--;
      $documentStore.config.line_options.splice(index, 1);
      $documentStore.blocks.forEach((block) => {
        block.columns.forEach((col) => {
          col.lines.splice(index, 1);
        });
      });
      $documentStore = $documentStore; // Trigger update
    }
  }
</script>

<div class="block-container">
  {#if $documentStore.config.show_labels}
    <div class="labels-column">
      {#each $documentStore.config.line_options as option, i}
        <div class="label-cell">
          <input
            type="checkbox"
            bind:checked={option.visible}
            title="Toggle Visibility"
          />
          <input bind:value={option.label} class="label-input" />
          <button
            class="line-delete-btn"
            on:click={() => deleteLine(i)}
            title="Delete Line">×</button
          >
        </div>
      {/each}
    </div>
  {/if}

  <div class="gloss-block">
    {#each block.columns as column, i}
      <div class="gloss-column">
        {#each column.lines as line, j}
          {#if $documentStore.config.line_options[j] && $documentStore.config.line_options[j].visible}
            <div
              class="gloss-cell line-{j}"
              contenteditable
              bind:textContent={column.lines[j]}
              placeholder={$documentStore.config.line_options[j].label || "..."}
              on:keydown={(e) => handleKeydown(e, i, j)}
            ></div>
          {/if}
        {/each}
      </div>
    {/each}
  </div>

  <button class="delete-btn" on:click={() => dispatch("delete")}>×</button>
</div>

<style>
  .block-container {
    position: relative;
    padding: 10px;
    padding: 10px;
    border: 1px solid var(--border);
    border-radius: 8px;
    background: var(--bg-panel);
    display: flex;
    gap: 10px;
  }

  .labels-column {
    display: flex;
    flex-direction: column;
    gap: 4px; /* Matched to gloss-column gap */
    padding-top: 0; /* Removed padding to align with first cell */
    min-width: 100px;
    border-right: 1px solid var(--border);
  }

  .label-cell {
    display: flex;
    align-items: center;
    gap: 5px;
    height: 32px; /* Fixed height to match gloss-cell (approx 1.5em + padding + border) */
    /* gloss-cell: min-height 1.5em (~24px) + 8px padding + 2px border = ~34px. Let's be explicit. */
  }

  .gloss-column {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 80px;
  }

  .gloss-cell {
    border: 1px solid transparent;
    padding: 4px 8px;
    height: 32px; /* Fixed height for alignment */
    min-height: 32px;
    line-height: 1.5; /* Explicit line height */
    white-space: pre; /* Prevent wrapping to maintain height */
    overflow: hidden; /* Hide overflow */
    word-break: break-all;
    word-break: break-all;
    border-radius: 4px;
    background: var(--bg-input);
    color: var(--text-main);
    font-family: "Courier New", Courier, monospace;
    font-size: 0.9rem;
  }

  .gloss-cell:focus {
    outline: none;
    background-color: var(--bg-app);
    border-color: var(--primary);
  }

  .gloss-cell.source {
    font-weight: bold;
    color: var(--accent);
  }

  .gloss-cell.translation {
    font-style: italic;
    color: var(--text-muted);
  }

  .delete-btn {
    position: absolute;
    top: 5px;
    right: 5px;
    background: transparent;
    border: none;
    color: var(--text-muted);
    font-size: 1.2rem;
    cursor: pointer;
    padding: 0 5px;
    line-height: 1;
  }

  .delete-btn:hover {
    color: var(--danger);
  }
</style>
