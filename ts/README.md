
---
### Importantes

Uso
- Quando der erro de importacao é porque houve mudanças na lib, solucao >>
  - Apague seu arquivo `deno.lock` para acompanhar as mudanças - após o taks do deno ler 2 vezes o arquivo ele para o erro.
  - ErroComum: `Integrity check failed for remote specifier. The source code is invalid, as it does not match the expected hash in the lock file`
- Importar a libs >>
  -```
  export type { LibTS as ILibTS } from 'https://raw.githubusercontent.com/reyzzao/libsRZ/main/ts/Init.ts'
  export { LibTS as libTS } from 'https://raw.githubusercontent.com/reyzzao/libsRZ/main/ts/Init.ts'
  ```

---