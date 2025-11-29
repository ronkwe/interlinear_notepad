<script>
  import Toolbar from "./components/Toolbar.svelte";
  import GlossBlockEditor from "./components/GlossBlockEditor.svelte";
  import { documentStore } from "./stores.js";

  function addBlock() {
    const newBlock = {
      columns: [{ lines: Array($documentStore.config.line_count).fill("") }],
    };
    $documentStore.blocks = [...$documentStore.blocks, newBlock];
  }

  function deleteBlock(index) {
    $documentStore.blocks = $documentStore.blocks.filter((_, i) => i !== index);
  }
</script>

<main>
  <h1>Glossa</h1>
  <Toolbar />

  <div class="blocks-container">
    {#each $documentStore.blocks as block, i}
      <GlossBlockEditor
        bind:block={$documentStore.blocks[i]}
        on:delete={() => deleteBlock(i)}
      />
    {/each}
  </div>

  <button class="add-block-btn" on:click={addBlock}>+ Add Block</button>
</main>

<style>
  main {
    text-align: left;
    padding: 20px;
    max-width: 1200px;
    margin: 0 auto;
  }

  .blocks-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    margin-bottom: 20px;
  }

  .add-block-btn {
    display: block;
    width: 100%;
    padding: 10px;
    background-color: var(--bg-panel);
    border: 2px dashed var(--border);
    color: var(--text-muted);
    cursor: pointer;
    border-radius: 8px;
    transition: all 0.2s;
  }

  .add-block-btn:hover {
    background-color: var(--bg-input);
    border-color: var(--primary);
    color: var(--primary);
  }
</style>
