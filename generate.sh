cd web/front
yarn
yarn build
cd ../admin
yarn
yarn build
cd ../../
mkdir ./www
cp -r ./web/front/dist/* ./www
cp -r ./web/admin/dist/* ./www
cp -r ./static ./www
cp ./static/*.* ./www
