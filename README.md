![title3](https://user-images.githubusercontent.com/24839815/130515256-f7fbebf9-f214-42f3-9cda-856458787f48.PNG)

# HashLion
HashLion is a basic (currently basic) password recovery utility coded in Go

## Installation
HashLion is exremely easy to install, you can download the latest release or clone the
repository and open it in vscode, then run in the vscode terminal the ```go build``` command, 
in both cases you'll end up with an executable file that you can just double click to start

## Usage
### Step 1 (providing a dictionary)
After double clicking the file you'll get a command prompt window asking you to enter
the path to your words file, the words list must be a .txt file

The current version of HashLion can only crack hashes by dictionary attacking
them so naturally the bigger your words list is, the better results you'll get
as you'll be able to crack more hashes.

some words lists:
https://github.com/praetorian-inc/Hob0Rules/tree/master/wordlists

I personally recommend using an updated version of the RockYou list:
https://github.com/ohmybahgosh/RockYou2021.txt

<details>
  <summary>picture</summary>
  
![step1](https://user-images.githubusercontent.com/24839815/130453626-92c4e318-3856-483e-a2d7-2f28a75e0074.PNG)
  
</details>
  
### Step 2 (choosing a hash type)
After providing a dictionary you need to choose the type
of hash algorithm that your password target was hashed with
HashLion supports following hash types for recovery:
- sha1
- sha256
- sha512
- md5

Any other hashes are not supported and won't allow you
to move to the next step

<details>
  <summary>picture</summary>
  
![Step2](https://user-images.githubusercontent.com/24839815/130454253-f255366b-7532-45fa-ab8d-6e538afef415.PNG)
  
</details>

### Step 3 (input target & attack)
Its time to provide our tagret and start the attack
we simple copy-paste our hash to the program.

After pressing enter HashLion will ask you if you want
want get live information output of the attack, while seeing
the live information looks impressive and cool, I recommend
to choose the "hide" always as its WAY faster since HashLion
can focus on only preforming instructions that are related to
the attack itself

Tip! attack mode is set to hide by default, so you can choose it by pressing enter and not writing anything

### Output mods
<details>
  <summary>HIDE</summary>

```
HIDE

Performs smallest amount of actions to go through the dictionary as
fast as possible, but provides a smaller amount of data about the attack
```
![Step5hide](https://user-images.githubusercontent.com/24839815/130457717-3437e007-00e3-4660-a49c-0ee9980e8cdd.PNG)

  </details>

<details>
  <summary>SHOW</summary>

```
SHOW

Provides live data about the attack, details every attempt and counts the
amount of attempts
```

![Step5](https://user-images.githubusercontent.com/24839815/130457573-8042ea5a-b481-453d-9f35-51e3f29b6b1a.PNG)

</details>

## ToDo list
These are some features that will be coming in the future as the project progresses
you can also feel free to contirbute and help adding these

- New attack option BruteForce

- Option to use Goroutines to go through the dictionary faster by splitting the dictionary and letting each routine go through one of the parts at the same time

- Support more hashing algorithms

- Add a config file where you can configure the input questions automatically instead of answering them everytime that you open the program

- Add an option to make HashLion save the attack information into a text file if an attack is successful

