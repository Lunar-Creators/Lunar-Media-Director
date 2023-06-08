import os
import subprocess
import runpy
import time

def loading():
    print('\033[3m\033[33m{}\033[0m'.format('Loading...'))
# LOADING AFTER DEFINING FUNCTION
loading()

def wronginput():
    print('\033[1m\033[91m{}\033[0m'.format('INCORRECT INPUT'))
    time.sleep(1)
    os.system('cls' if os.name == 'nt' else 'clear')

def leavemodule():
    print('\033[3m\033[33m{}\033[0m'.format('Leaving...'))

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
    if choice == 'H' or choice == 'h':
        runpy.run_module(mod_name='ffmpeg-helpmodule')
        welcome()
    if choice == 'Y' or choice == 'y':
        void()
    if choice == 'A' or choice == 'a':
        void()
    if choice == 'P' or choice == 'p':
        void()
    if choice == 'T' or choice == 't':
        void()
    if choice == 'Q' or choice == 'q':
        void()
    if choice == 'K' or choice == 'k':
        void()
    if choice == 'C' or choice == 'c':
        os.system('cls' if os.name == 'nt' else 'clear')
        print('In this mode you can run your command with your flags for ffmpeg\nUse this mode if you fully know what you are doing. "-h" for help.\nType "quit" to exit\n\nEnter ffmpeg arguments directly,\nexample: "','\033[32m{}\033[0m'.format('-i input -c:v libvpx -crf 16 -an -lag-in-frames 16 -auto-alt-ref 0 -y output'),'"')
        def supercustom():
            cmd = input("ffmpeg ")
            if cmd == 'quit':
                welcome()
            else:
                subprocess.run(f'ffmpeg {cmd}', check=False)
                supercustom()
        supercustom()
    if choice == 'X' or choice == 'x':
        subprocess.run('explorer.exe https://github.com/Lunar-Creators/Lunar-Media-Director', check=False)
        welcome()
    if choice == 'V' or choice == 'v':
        void()
    if choice == 'E' or choice == 'e':
        leavemodule()
    else:
        wronginput()
        welcome()
# NEED TO CLEAR TERMINAL BEFORE RUNNING A COMMAND
os.system('cls' if os.name == 'nt' else 'clear')
welcome()
