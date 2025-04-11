import { workspace, ExtensionContext } from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  const platform = process.platform;

  let binaryName = "plantuml-lsp";
  if (platform === "win32") {
    binaryName += ".exe";
  } else if (platform === "darwin") {
    binaryName += "-darwin";
  }

  const config = workspace.getConfiguration("plantuml-lsp");
  // Read user-defined settings
  const stdlibPath = config.get<string>("stdlibPath", "");
  const execPath = config.get<string>("execPath", "");
  const jarPath = config.get<string>("jarPath", "");

  // Build the CLI arguments based on user settings
  let args: string[] = [];
  if (stdlibPath) args.push("--stdlib-path", stdlibPath);
  if (execPath) args.push("--exec-path", execPath);
  if (jarPath) args.push("--jar-path", jarPath);

  const serverModule = context.asAbsolutePath(binaryName);

  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used
  const serverOptions: ServerOptions = {
    run: { command: serverModule, transport: TransportKind.stdio, args: args },
    debug: {
      command: serverModule,
      transport: TransportKind.stdio,
      args: args,
    },
  };

  // Options to control the language client
  const clientOptions: LanguageClientOptions = {
    // Register the server for PlantUML documents
    documentSelector: [{ scheme: "file", language: "plantuml" }],
    synchronize: {
      // Notify the server about file changes to '.clientrc files contained in the workspace
      fileEvents: workspace.createFileSystemWatcher("**/.clientrc"),
    },
  };

  // Create the language client and start the client.
  client = new LanguageClient(
    "plantuml-lsp",
    "PlantUML LSP",
    serverOptions,
    clientOptions,
  );

  // Start the client. This will also launch the server
  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
