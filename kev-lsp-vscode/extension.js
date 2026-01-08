const { LanguageClient, TransportKind } = require('vscode-languageclient/node');

let client;

function activate(context) {
    const serverOptions = {
        command: '/Users/gagemcguire/Documents/GitHub/kev-lsp/kev-lsp',
        transport: TransportKind.stdio
    };

    const clientOptions = {
        documentSelector: [{ scheme: 'file', language: 'kev' }]
    };

    client = new LanguageClient('kev-lsp', 'KEV LSP', serverOptions, clientOptions);
    client.start();
}

function deactivate() {
    if (client) {
        return client.stop();
    }
}

module.exports = { activate, deactivate };