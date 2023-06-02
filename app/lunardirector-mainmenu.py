print('\033[3m\033[33m{}\033[0m'.format('Loading...'))

import os
import subprocess
import runpy
import time

def reset():
    del filepath 
    del inputaudio 
    del encoder 
    del outputformat 
    del outputname 
    del predel 
    del Profile 
    del tune 
    del vidbitrate 
    del audiocodec 
    del threads 
    del audiobitrate 
    del flags 
    del subencoder 
    del outputfolder 
    del inputsubtitle 
    del disablevideo 
    del disableaudio 
    del framerate 
    del size 
    del volume 
    del disablesubtitles 
    del audiotype 
    del maxbitrate 
    del choice

# exec = 'ffmpeg'
# optlist = f' {exec} input'
# ffmpeg = exec + optlist
# print(ffmpeg)
# opts = '-i F:\intro720.mov -c:v copy F:\intpython.mov'
# subprocess.run(f'ffmpeg {opts}')

def welcome():
    os.system('cls' if os.name == 'nt' else 'clear')
    print('\033[3m\033[33m{}\033[0m'.format('--appver unknown -indev'),'\033[3m\033[91m{}\033[0m'.format('--copyright "Lunar Creators"\n'))
    print('If you have a problem, or you want to suggest a preset or feature to add, please contact us!\n','\033[1m{}\033[0m'.format('WELCOME TO LUNAR MEDIA DIRECTOR! (CLI Version)'),'\n \n-------------------------- \n  H - FFmpeg Help\n  Y - Select Video Preset\n  A - Audio Converting\n  P - Photo Converting\n  T - Select Tool\n  Q - Half-Manual Mode\n  K - Audio to Video Encoding\n  C - Commandline mode\n  X - Our Github\n  V - Open Video Downloader\n  E - Exit \n--------------------------')
    choice = input("Type your choice: ")
    if choice == 'H':
        runpy.run_module(mod_name='ffmpeg-helpmodule')
        welcome()
    if choice == 'Y':
        void()
    if choice == 'A':
        void()
    if choice == 'P':
        void()
    if choice == 'T':
        void()
    if choice == 'Q':
        void()
    if choice == 'K':
        void()
    if choice == 'C':
        os.system('cls' if os.name == 'nt' else 'clear')
        print('In this mode you can run your command with your flags for ffmpeg\nUse this mode if you fully know what you are doing. "-h" for help.\nType "quit" to exit\n\nEnter ffmpeg arguments directly,\nexample: "','\033[32m{}\033[0m'.format('-i input -c:v libvpx -crf 16 -an -lag-in-frames 16 -auto-alt-ref 0 -y output'),'"')
        def supercustom():
            cmd = input("ffmpeg ")
            if cmd == 'quit':
                welcome()
            else:
                subprocess.run(f'ffmpeg {cmd}')
                supercustom()
        supercustom()
    if choice == 'X':
    if choice == 'X':
        subprocess.run('explorer.exe https://github.com/Lunar-Creators/Lunar-Media-Director')
        welcome()
    if choice == 'V':
        void()
    if choice == 'E':
        void()
    else:
        print('\033[1m\033[91m{}\033[0m'.format('INCORRECT INPUT'))
        time.sleep(1)
        os.system('cls' if os.name == 'nt' else 'clear')
        welcome()
welcome()

