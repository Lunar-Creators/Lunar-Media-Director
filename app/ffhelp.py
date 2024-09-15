import os
import subprocess
import time
import ldmenu
import lunlib

lunlib.loading()
    
def helpmenu():
    lunlib.clear()
    print('\033[1m{}\033[0m'.format('\nHow can ffmpeg help you?'),'\n--------------------------\n  1 - Standart list of help\n  2 - Very long List of help\n  3 - EXTREAMLY LONG LIST OF HELP\n  4 - Flag help for available ffmpeg encoders\n  5 - List of available ffmpeg encoders\n  Q - Quit Help\n--------------------------\n','\033[35m{}\033[0m'.format('NOTE! In commandline mode, you can enter the help command with your arguments'))
    choice = input("Type your choice: ")
    if choice == '1':
        subprocess.run('ffmpeg -h', check=False)
        helpmenu()
    if choice == '2':
        subprocess.run('ffmpeg -h long', check=False)
        helpmenu()
    if choice == '3':
        subprocess.run('ffmpeg -h full', check=False)
        helpmenu()
    if choice == '4':
        subprocess.run('ffmpeg -encoders', check=False)
        print('\033[1m{}\033[0m'.format('\n\nEnter the one you are interested in from the list'),'\033[33m{}\033[0m'.format('\n(example: libvpx-vp9)'),)
        encoder = input("You need help with: ")
        subprocess.run(f'ffmpeg -h encoder={encoder}', check=False)
        del encoder
        helpmenu()
    if choice == '5':
        subprocess.run('ffmpeg -encoders', check=False)
        helpmenu()
    if choice == 'Q':
        ldmenu.welcome()
    else:
        lunlib.wronginput()
        helpmenu()
lunlib.clear()
helpmenu()    