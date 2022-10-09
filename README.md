# Scraper GoldenManga

- [x] Pega todas as rotas da grade /mangas
- [x] Acessa a pagina deles pela rota /mangabr
    - [x] Pega o titulo
    - [x] Pega a categoria
    - [x] Pega o status de lançamento
    - [x] Pega o autor
    - [x] Pega o artista
    - [x] Pega a descrição
    - [ ] Pega a capa 
    - Capitulos
        - [x] Pega a data dos capitulos
        - [x] Pega a imagem das paginas

## Coisas a serem adicionadas depois de terminar o Scraper Principal

- Adicionar a opção de buscar (acredito que ficara mais rapida)
- Pegar os mangas em destaques
- Pegar as ultimas atualizações
- Pegar manga por categoria
- Adicionar mais novidades(Procrastinação :D)

#### Se for utilizar apenas o scraper para algo a mais
---
> Aqui vai apenas a função que faz tudo, ou
você pode fuçar o codigo e fazer alterações que desejar 
```
# Sintax da função

func StartScraper(
    ScraperPerPage int, 
    OutFileName string
)
```