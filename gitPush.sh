git add .

echo "Masukans Pesan Commit : "
read message

git commit -m "$message"

git push origin master 

echo "Done!"