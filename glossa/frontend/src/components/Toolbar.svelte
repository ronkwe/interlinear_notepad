<script>
  import { OpenFile, SaveFile, ExportToMarkdown, ExportToPDF } from '../../wailsjs/go/main/App'

  let doc = {} // This would be the main document state

  async function openFile() {
    try {
      const result = await OpenFile()
      if (result) {
        doc = result
        // TODO: Update the app's main document state
        alert('File opened successfully!')
      }
    } catch (err) {
      console.error(err)
      alert('Error opening file: ' + err)
    }
  }

  async function saveFile() {
    try {
      const path = await SaveFile(doc)
      if (path) {
        alert('File saved successfully to: ' + path)
      }
    } catch (err) {
      console.error(err)
      alert('Error saving file: ' + err)
    }
  }

  async function exportMarkdown() {
    try {
      const markdown = await ExportToMarkdown(doc)
      // This would typically be saved to a file, but for now we'll just show it
      console.log(markdown)
      alert('Markdown export successful! Check the console.')
    } catch (err) {
      console.error(err)
      alert('Error exporting to Markdown: ' + err)
    }
  }

  function exportPDF() {
    try {
      ExportToPDF()
    } catch (err) {
      console.error(err)
      alert('Error exporting to PDF: ' + err)
    }
  }
</script>

<div class="toolbar">
  <button on:click={openFile}>Open</button>
  <button on:click={saveFile}>Save</button>
  <button on:click={exportMarkdown}>Export Markdown</button>
  <button on:click={exportPDF}>Export PDF</button>
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
    border: 1px solid transparent;
    font-weight: 500;
    font-family: inherit;
    background-color: #1a1a1a;
    cursor: pointer;
    transition: border-color 0.25s;
  }
  button:hover {
    border-color: #a855f7;
  }
  button:focus,
  button:focus-visible {
    outline: 4px auto -webkit-focus-ring-color;
  }
</style>
