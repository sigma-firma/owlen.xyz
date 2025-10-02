# Owlen - Arxiv Catalogue
Owlen is a demo web app which catalogues the Arxiv database contained in the 
JSON file they provide at 
https://www.kaggle.com/datasets/Cornell-University/arxiv

Owlen allows users to filter research papers by tags, which are parsed from the
data contained in the dataset. 

# Functionality
 - Filter by tag/subject
 - View by tag/subject
 - Follow tags

This is only the web app part of the program. The other part, which parses the
dataset, is not published on github, because I'm always tweaking it. 

Owlen is built using the following:
 - Golang is used on the backend/server code, as well as the parser. 
 - JavaScript/ajax/html/css is used client side.
 - Redis is used to store and retrieve data.
