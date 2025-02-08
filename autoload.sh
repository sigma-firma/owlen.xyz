# use in n/vim to restart on save:
# :autocmd BufWritePost * silent! !./autoload.sh
#!/bin/bash
pkill owlen.xyz || true
go build -o owlen.xyz
echo http://localhost:10528
./owlen.xyz >> log.txt 2>&1 &
