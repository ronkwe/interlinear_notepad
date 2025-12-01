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

<div class="legend-container">
    <div class="legend-header">
        <h3>Legend</h3>
        <button class="add-line-btn" on:click={addLine}>+ Line</button>
    </div>
    <div class="legend-lines">
        {#each $documentStore.config.line_options as option, i}
            <div class="legend-row">
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
</div>

<style>
    .legend-container {
        background: var(--bg-panel);
        border: 1px solid var(--border);
        border-radius: 8px;
        padding: 15px;
        margin-bottom: 20px;
        width: fit-content;
        min-width: 250px;
    }

    .legend-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 10px;
    }

    .legend-header h3 {
        margin: 0;
        font-size: 1.1rem;
        color: var(--text-main);
    }

    .add-line-btn {
        background: var(--primary);
        border: 1px solid var(--primary);
        color: var(--bg-app);
        padding: 6px 12px;
        border-radius: 4px;
        cursor: pointer;
        font-size: 0.9rem;
        font-weight: bold;
        transition: opacity 0.2s;
    }

    .add-line-btn:hover {
        opacity: 0.9;
    }

    .legend-lines {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .legend-row {
        display: flex;
        align-items: center;
        gap: 8px;
    }

    .label-input {
        background: var(--bg-input);
        border: 1px solid var(--border);
        color: var(--text-main);
        padding: 4px 8px;
        border-radius: 4px;
        flex-grow: 1;
    }

    .line-delete-btn {
        background: transparent;
        border: none;
        color: var(--text-muted);
        font-size: 1.2rem;
        cursor: pointer;
        padding: 0 4px;
        line-height: 1;
    }

    .line-delete-btn:hover {
        color: var(--danger);
    }
</style>
