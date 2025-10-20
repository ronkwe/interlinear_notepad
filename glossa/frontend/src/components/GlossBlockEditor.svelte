<script>
  export let block

  function handleKeydown(event, columnIndex) {
    const el = event.target
    const selection = window.getSelection()

    // Add a new column when space is pressed at the end of a cell
    if (event.key === ' ' && selection.anchorOffset === el.innerText.length) {
      event.preventDefault()
      insertColumnAfter(columnIndex)
    }

    // Merge with previous column if backspace is pressed at the start of a cell
    if (event.key === 'Backspace' && selection.anchorOffset === 0 && columnIndex > 0) {
      event.preventDefault()
      mergeWithPreviousColumn(columnIndex)
    }
  }

  function insertColumnAfter(index) {
    const newColumn = { lines: Array(block.columns[0].lines.length).fill('') }
    block.columns.splice(index + 1, 0, newColumn)
    block.columns = block.columns // Trigger Svelte reactivity
    // TODO: Focus on the first line of the new column
  }

  function mergeWithPreviousColumn(index) {
    const currentColumn = block.columns[index]
    const previousColumn = block.columns[index - 1]

    currentColumn.lines.forEach((line, i) => {
      previousColumn.lines[i] += line
    })

    block.columns.splice(index, 1)
    block.columns = block.columns // Trigger Svelte reactivity
    // TODO: Set cursor position correctly in the merged cell
  }
</script>

<div class="gloss-block-editor">
  {#each block.columns as column, columnIndex}
    <div class="gloss-column">
      {#each column.lines as line, lineIndex}
        <div
          class="gloss-cell"
          contenteditable="true"
          bind:innerHTML={column.lines[lineIndex]}
          on:keydown={(e) => handleKeydown(e, columnIndex)}
        ></div>
      {/each}
    </div>
  {/each}
</div>

<style>
  .gloss-block-editor {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(80px, max-content));
    gap: 5px;
    border: 1px solid #555;
    padding: 10px;
    margin-top: 20px;
    overflow-x: auto;
  }

  .gloss-column {
    display: contents;
  }

  .gloss-cell {
    border: 1px solid #444;
    padding: 8px;
    min-height: 2em;
    white-space: pre-wrap;
    word-break: break-all;
  }

  .gloss-cell:focus {
    outline: 2px solid #a855f7;
    background-color: #3a3a3a;
  }
</style>
