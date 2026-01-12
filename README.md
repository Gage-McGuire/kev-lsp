### KEV's Language Server

`./kev-lsp-vscode-extension` contains the files to run the LSP in vscode  
`./rpc` contains code pertaining to the communication protocol (encoding/decoding)   
`./lsp` contains mostly models/structs which allows for marshaling and unmarshaling  
`./analysis` contains logic and states which determain responses   
`./handler` contains logic to route/handle requests/responses down correct paths  
`./logger` contains code that logs lsp data to defined log file
