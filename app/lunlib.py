def inputfile():
    from tkinter import filedialog as file

    filetypes = (
            ('Text Files', '*.txt'),
            ('All files', '*.*')
        )
    fileselect = file.askopenfilename(
            title='Select file',
            initialdir='/',
            filetypes=filetypes)
    print(f'{fileselect}')
    
# lunlib.inputfile()
    
def loading():
    print('\033[3m\033[33m{}\033[0m'.format('Loading...'))

# lunlib.loading()

def wronginput():
    import os
    import time
    import lunlib
    print('\033[1m\033[91m{}\033[0m'.format('INCORRECT INPUT'))
    time.sleep(1)
    lunlib.clear()
    
# lunlib.wronginput()

def leavemsg():
    print('\033[3m\033[33m{}\033[0m'.format('Leaving...'))
    
# lunlib.leavemsg()

def clear():
    import os
    os.system('cls' if os.name == 'nt' else 'clear')
    
# lunlib.clear()