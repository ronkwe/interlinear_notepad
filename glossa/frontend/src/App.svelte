<script>
  import Toolbar from './components/Toolbar.svelte'
  import GlossBlockEditor from './components/GlossBlockEditor.svelte'

  let doc = {
    version: '1.0',
    config: {
      line_count: 4,
      line_labels: ['Source', 'Morphemes', 'Gloss', 'Translation'],
    },
    blocks: [
      {
        columns: [
          { lines: ['Wa\'khnínu\'s', 'wa\'-k-hnínu-\'-s', 'FACT-1S.PAT-buy-PUNC-PL', 'I bought them.'] },
          { lines: ['pésku', 'pésku', 'cat', 'the cat'] },
        ],
      },
    ],
  }
</script>

<main>
  <h1>Glossa</h1>
  <Toolbar />
  <GlossBlockEditor bind:block={doc.blocks[0]} />
</main>

<style>
  main {
    text-align: left;
  }
</style>
