
---
### Importantes

Uso
- Quando der erro de importacao é porque houve mudanças na lib, solucao >>
  - Apague seu arquivo `deno.lock` para acompanhar as mudanças - após o taks do deno ler 2 vezes o arquivo ele para o erro.
  - ErroComum: `Integrity check failed for remote specifier. The source code is invalid, as it does not match the expected hash in the lock file`
- Importar a libs >>
  -```
  // uso: AUTO_ATUALIZACOES (se a lib mudar muda junto), exclua-o lock.json, e importe a lib
import { libTS } from 'https://raw.githubusercontent.com/reyzzao/libsRZ/main/ts/Init.ts'
// usando a lib importada em polo unico com essas vars
export const Log = libTS.Log

  ```

---