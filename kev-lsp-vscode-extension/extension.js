const { LanguageClient, TransportKind } = require('vscode-languageclient/node');
require('dotenv').config({ path: __dirname + '/../.env' });

let client;

function activate(context) {
    const serverOptions = {
        command: process.env.LSP_PATH,
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