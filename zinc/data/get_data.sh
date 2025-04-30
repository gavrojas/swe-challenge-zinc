#!/bin/bash
# Download the file 
URL_FILE="http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz"
echo "Downloading file"
wget -c $URL_FILE

# name of file .tgz - contains the data of the mails
TAR_FILE="enron_mail_20110402.tgz"

# decompress de file in the directory
echo "Decompressing the file $TAR_FILE ..."
tar -zxf "$TAR_FILE" -C .

echo "Decompression complete. Files extracted in the directory."
