import { writable } from 'svelte/store';

export const documentStore = writable({
    version: '1.0',
    config: {
        line_count: 4,
        line_options: [
            { label: 'Source', visible: true },
            { label: 'Morphemes', visible: true },
            { label: 'Gloss', visible: true },
            { label: 'Translation', visible: true },
        ],
        show_labels: true,
    },
    blocks: [
        {
            columns: [
                { lines: ['', '', '', ''] },
            ],
        },
    ],
});
