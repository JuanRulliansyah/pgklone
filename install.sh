#!/bin/sh

repo_url="https://github.com/JuanRulliansyah/pgklone/releases/download/v1.0.0/pgklone"
install_dir="/usr/local/bin"
target_name="pgklone"

echo "Downloading the latest version of $target_name from $repo_url..."
curl -L "$repo_url" -o "/tmp/$target_name"

if [ ! -f "/tmp/$target_name" ]; then
  echo "Download failed!"
  exit 1
fi

echo "Installing $target_name to $install_dir..."
sudo mv "/tmp/$target_name" "$install_dir/$target_name"

echo "Making $target_name executable..."
sudo chmod +x "$install_dir/$target_name"

if [ -x "$install_dir/$target_name" ]; then
  echo "$target_name installed successfully!"
else
  echo "Failed to install $target_name"
  exit 1
fi
