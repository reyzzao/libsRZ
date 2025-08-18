
---
### Explicacoes

### Usando

async function mainModules() {
  console.log('Iniciando a app mainModules...')

  // Define o caminho base onde os módulos estão localizados.
  const modulesBasePath = new URL('./modules/', import.meta.url).href

  const allModules = await loaderModulesIndex(modulesBasePath, [
    'example_foo',
    // ... adicione aqui o nome das suas outras pastas que têm Index.ts
  ])
  // console.dir(allModules);
  return allModules
}

const useMainModules = await mainModules()
Log(useMainModules.foo)

---