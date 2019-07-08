const go = new Go();

WebAssembly.instantiateStreaming(fetch("formatter.wasm"), go.importObject).then(
  result => {
    go.run(result.instance);
  }
);
