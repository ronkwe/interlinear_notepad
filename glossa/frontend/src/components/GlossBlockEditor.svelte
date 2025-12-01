<script>
  import { createEventDispatcher, onMount, onDestroy, tick } from "svelte";
  import { documentStore } from "../stores.js";
  export let block;

  const dispatch = createEventDispatcher();

  function handleKeydown(event, columnIndex, lineIndex) {
    const el = event.target;
    const selection = window.getSelection();

    // Add a new column when space is pressed at the end of a cell
    if (event.key === " " && selection.anchorOffset === el.innerText.length) {
      event.preventDefault();
      if (columnIndex + 1 < block.columns.length) {
        // Navigate to next column
        const blockEl = el.closest(".gloss-block");
        const columns = blockEl.querySelectorAll(".gloss-column");
        const nextColumn = columns[columnIndex + 1];
        if (nextColumn) {
          const cells = nextColumn.querySelectorAll(".gloss-cell");
          // Try to focus the same line index, fallback to first line
          const targetCell = cells[lineIndex] || cells[0];
          if (targetCell) targetCell.focus();
        }
      } else {
        insertColumnAfter(columnIndex);
      }
    }

    // Add a spacer and new column when '.' is pressed at the end of a cell
    if (event.key === "." && selection.anchorOffset === el.innerText.length) {
      // Allow the '.' to be typed first, then trigger the logic
      // We use setTimeout to let the browser update the text content
      setTimeout(() => {
        insertSpacerAndNewColumn(columnIndex);
      }, 0);
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

  function generateUUID() {
    return crypto.randomUUID();
  }

  function insertColumnAfter(index) {
    const newColumn = {
      id: generateUUID(),
      lines: Array($documentStore.config.line_count).fill(""),
    };
    block.columns.splice(index + 1, 0, newColumn);
    block.columns = block.columns; // Trigger Svelte reactivity
    // TODO: Focus on the first line of the new column
  }

  function insertSpacerAndNewColumn(index) {
    // 1. Insert Spacer Column (Empty)
    const spacerColumn = {
      id: generateUUID(),
      lines: Array($documentStore.config.line_count).fill(""),
      isSpacer: true,
    };

    // 2. Insert New Column for next sentence
    const newColumn = {
      id: generateUUID(),
      lines: Array($documentStore.config.line_count).fill(""),
    };

    block.columns.splice(index + 1, 0, spacerColumn, newColumn);
    block.columns = block.columns; // Trigger Svelte reactivity

    // Focus on the new column (index + 2)
    tick().then(() => {
      const blockEl = container
        .querySelector(".gloss-block-row")
        .closest(".block-container");
      // We need to find the correct column element. Since we have rows now, it's a bit trickier.
      // But we can query all gloss-columns in the container and find the one at the new index.
      // Wait, the index in `block.columns` matches the flattened list of columns.
      // But `querySelectorAll(".gloss-column")` will return them in order.
      const allColumns = blockEl.querySelectorAll(".gloss-column");
      const targetColumn = allColumns[index + 2];
      if (targetColumn) {
        const firstCell = targetColumn.querySelector(".gloss-cell");
        if (firstCell) firstCell.focus();
      }
    });
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

  let container;
  let containerWidth = 0;
  let resizeObserver;
  let groupedColumns = [];

  $: {
    if (block.columns && containerWidth) {
      calculateLayout();
    }
  }

  function calculateLayout() {
    if (!containerWidth) return;

    const labelWidth = 100; // Min width of label column
    const columnWidth = 84; // Min width (80) + gap (4)
    const availableWidth = containerWidth - labelWidth - 20; // -20 for padding
    const columnsPerRow = Math.max(1, Math.floor(availableWidth / columnWidth));

    groupedColumns = [];
    const totalItems = block.columns.length + 1; // +1 for the "Add Sentence" button

    for (let i = 0; i < totalItems; i += columnsPerRow) {
      const groupCols = block.columns.slice(i, i + columnsPerRow);
      // Check if this group should contain the button
      // The button is the last item (index totalItems - 1)
      // If the range [i, i + columnsPerRow) includes totalItems - 1, then yes.
      const hasButton = i + columnsPerRow >= totalItems;

      groupedColumns.push({
        startIndex: i,
        columns: groupCols,
        hasButton: hasButton,
      });
    }
  }

  onMount(() => {
    resizeObserver = new ResizeObserver((entries) => {
      for (let entry of entries) {
        containerWidth = entry.contentRect.width;
      }
    });
    if (container) resizeObserver.observe(container);
  });

  onDestroy(() => {
    if (resizeObserver) resizeObserver.disconnect();
  });
</script>

<div class="block-container" bind:this={container}>
  {#if groupedColumns.length > 0}
    {#each groupedColumns as group, groupIndex}
      {#if groupIndex > 0}
        <hr class="row-separator" />
      {/if}
      <div class="gloss-row">
        {#if $documentStore.config.show_labels}
          <div class="labels-column">
            {#each $documentStore.config.line_options as option, i}
              <div class="label-cell">
                <span class="truncated-label"
                  >{option.label.substring(0, 3)}</span
                >
              </div>
            {/each}
          </div>
        {/if}

        <div class="gloss-block-row">
          {#each group.columns as column, i (column.id)}
            <div class="gloss-column">
              {#each column.lines as line, j}
                {#if $documentStore.config.line_options[j] && $documentStore.config.line_options[j].visible}
                  {#if column.isSpacer}
                    <div class="gloss-cell spacer-cell"></div>
                  {:else}
                    <div
                      class="gloss-cell line-{j}"
                      role="textbox"
                      tabindex="0"
                      contenteditable
                      bind:textContent={column.lines[j]}
                      placeholder={$documentStore.config.line_options[j]
                        .label || "..."}
                      on:keydown={(e) =>
                        handleKeydown(e, group.startIndex + i, j)}
                    ></div>
                  {/if}
                {/if}
              {/each}
            </div>
          {/each}
          {#if group.hasButton}
            <div class="gloss-column add-sentence-column">
              <button
                class="add-sentence-btn"
                title="Add New Sentence"
                on:click={() =>
                  insertSpacerAndNewColumn(block.columns.length - 1)}>+</button
              >
            </div>
          {/if}
        </div>
      </div>
    {/each}
  {:else}
    <!-- Fallback/Initial render -->
    <div class="gloss-row">
      <!-- ... (simplified fallback if needed, or just let it be empty until resize triggers) ... -->
    </div>
  {/if}

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
    flex-direction: column;
    gap: 10px;
  }

  .gloss-row {
    display: flex;
    gap: 10px;
    align-items: flex-start;
  }

  .row-separator {
    width: 100%;
    border: 0;
    border-top: 1px solid var(--border);
    margin: 5px 0;
  }

  .truncated-label {
    font-size: 0.8rem;
    color: var(--text-muted);
    font-weight: bold;
    text-transform: uppercase;
    padding-left: 5px;
  }

  .gloss-block-row {
    display: flex;
    gap: 4px;
    flex-wrap: nowrap; /* No wrapping within a row */
  }

  .label-cell {
    display: flex;
    align-items: center;
    gap: 5px;
    height: 32px; /* Fixed height to match gloss-cell */
    padding: 4px 8px; /* Match gloss-cell padding */
    border: 1px solid transparent; /* Match gloss-cell border */
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

  .gloss-cell.spacer-cell {
    background: var(--bg-panel);
    border: 1px dashed var(--border);
    pointer-events: none;
    opacity: 0.5;
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

  .add-sentence-column {
    justify-content: center;
    min-width: 40px; /* Smaller than regular column */
  }

  .add-sentence-btn {
    width: 100%;
    height: 100%;
    background: transparent;
    border: 1px dashed var(--border);
    color: var(--text-muted);
    font-size: 1.5rem;
    cursor: pointer;
    border-radius: 4px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s;
  }

  .add-sentence-btn:hover {
    border-color: var(--primary);
    color: var(--primary);
    background: var(--bg-input);
  }
</style>
