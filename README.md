#### CUDA Version Switcher

I made this script to switch the version of CUDA for fish terminal. It can detect the versions of CUDA installed and it will list it for you to pick one.

Why?
My current project uses multiple models that requires different CUDA version. I want to train my model on CUDA 12.3, but I need to preprocess my data (images) with [Megadetector](https://github.com/microsoft/CameraTraps/blob/main/megadetector.md) that uses CUDA 11.3 and python 3.8. Switching it manually takes a minute I know, but I also wanted to make something with GoLang.


- It only works with fish terminal
- It assumes CUDA versions are saved in `/usr/local`
- It assumes fish config is saved in `$HOME/.config/fish/config.fish`
- Also it only works with linux

  ## Install

  Just add /bin to path and run `CUDAswitch`

  
![image](https://github.com/hamoodev/CUDAswitcher/assets/45039279/cdd08826-6746-43b7-aef8-79f6378abd07)
