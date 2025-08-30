# LifeSync
A calendar that optimises your time. 

Will be written in golang, using fyne. 

Goals:

Make a GUI that shows calendar, (week, month, year), sort of like google calendar. 

Backend: 

get dates working, so sort of like, make an year class, then month, then day, represented as a numerical value. 

year -> months -> days 


This project has been neglected for a while, will do the to do functionality. 

go 


To do app part:

Need to make a completed task method, this marks the task as completed. need to add this as a parameter to the table. 


need to add priority as a parameter to the table. then need to have functions that utilise three different parameters, high, medium, and low. 


Then need to add the UI elements, so a text field for inputting tasks, 3 buttons to choose the priority option, then also perhaps an estimated amount of time taken to it.

For now this will work as a basic model: 

Extensions: 

In the future edit the still active tasks, so say you've spent 30 minutes on a task, we want to deduct that from the estimated total time, and add this to the time spent on it category. This lets you keep track of time spent on the task, Vs how much you estimated it will take. 


Want to have different sections, based on priority of tasks, so high priority tasks are shown first, medium second etc. Also want to combine this with deadlines, so say if a medium priority task has a deadline today that should be shown above a high priority task that's due in 1 month from now etc. Then add the estimated time for completion into this algorithm etc. 


Another extension would be to separate the tasks into various categories, or add a task category section. You can make new categories if needed. So one category can be "reading" and you can track how much time you spend reading books etc. You can then nest these so you have categories inside categories, so a second category such as "Catcher in the rye," so it's sort of like in this category you've spent X amount of time on this task. 


Add recurring tasks, so tasks you have to do every week, or every day etc. Ie your morning routine, or bedtime routine etc. 


Integrate points, so you get points for completion of tasks, sort of silly but I think is a nice addition.Once the points go above a certain amount you can go and do something fun or outside the routine. Draconic I know. 


Want to view completed tasks, and want to add a search engine based on the title of the task, this way you can save time instead of having to rewrite a previously completed task, I also want to add an "improvements," strategy, where you can add notes to this "task type" based on the title, on how you can save time in the future or whatever you want to note about this sort of task, perhaps ways to perform better at this activity etc, sort of a Journal for every specific activity where you can reflect and then look it up whenever you want, re-evaluate and improve upon. 


-----------------
Add contacts page,

initial contact page -- done 
basic description per contact -- done 

Similar to add tasks, but for people. So you can add new people you meet on here, their social media information, contact information, and a note about them, their birthdays, important days etc, Say they tell you they are leaving the country at X time etc, you can remember to say goodbye or smth, or buy a gift etc. This is again sort of like a task, but related to people instead. There's also a note section where you can take note about what they like sort of their favourite foods/colour etc. 


Also it can remind you to contact people every now and then so you stay in touch. Perhaps automatic messages like Happy Birthday or Happy Christmas etc. This saves time, and helps you keep track of things. You can then manually respond back to them if they say anything back. 

So this requires some API stuff I assume to make it automatically run in the background. 

Also reminds you to buy gifts for them, so this is based on whether you want to give them a gift or not, and if so it reminds you 2 weeks before, this should be a variable option, obviously more important people you should think about what to give etc. 


Should have a central pie chart showing how much time was spend on this "person" admin stuff. Should obviously try and automate as much of this as possible. 

Also add time spend on doing these tasks for these people. And then have a pie chart showing this much time spent this month on this person etc etc. Should also be able to change the time frame as to when you want to analyse this data etc. 

Also maybe compare this with the amount of time you think or estimate they spent on you and it'll be interesting to see the contrast heh.

----------------

Add a 





















---------------------------------------------------------------------
 So just focuing on the Task lists UI at the moment. 

I'm thinking of having some sort of scrollable, element, with active tasks at the start, and completed tasks at the end, or beind the active tasks, sorting can be figured out later, no 

so sorted by time, hours taken, points, etc. can pick how to sort, 

but the distinction between active and done tasks always exists, so active tasks first, inactive after, 

need a search function, or algorithm that can do this, so maybe custom write one????? meh 

Should be able to click on a task, to "expand" it, which provides more information about the task in question, do error handling here, in case it is empty, and the database structure needs to be 
changed yet again. 

Add a neat plus button at the very top of the scroll section, or maybe to the right hand side top of the screen to add a new task. 

Search bar can also be near the add button I suppose. 

usual we have been taught to adapt frontend to the backend, but really should be the opposite. 








__________________________________________________________________________________________________________________________



03/08/2025 

So tasks page:

needs to show add task button green  
needs to add a tasks description green 
need to add a point enter button green 
need to enter a point field, which is a number should be checked if it is a number or not (backend)
task should show time of completion/ est time to complete, or time left to finish est. green 
est time show in task green
priority show the priority of the task in the task section, as a colour? colour coded stich to high medium low for now green  



needs to be able to display the tasks
which you click on the task it needs to be expanded a bit,  green 
if you click on the expanded section's a button or somthing, like "read more" or "expand" it shuold then go to a new page link, 
expanded part should also have an option to edit the task, this is important need to add *
* 





So I finished the add button
finished the add tasks and add points (number)
div task list has a list of tasks and is meant to deal with the list of tasks and show it. 
also did some styling for the front page, need to do the hidden task overview thing next soon, asap



5/08/2025 

So next steps would be styling add button 
style inputs points and task description 

fix styling on the site, so choose a colour scheme and stick with it, dark mode / colour scheme
maybe black and purple colour scheme, add navy or blue as well to this, green for highlights 

06/08/2025

Chose the colour scheme. green 
So tidied up the webpage a bit, and tidied up the colour scheme. 

07/08/2025

Going to change to light theme ahh, probably easier to make it look nicer. Well it should be reativelty easier to make it look nicer. 
If it was just me I'd keep it dark and white ore green / blue. 



08/08/2025

Okay, so I'm going for a very simple colour dark scheme, inspired by google calendar. 




DARK COLOUR SCHEME:

Background: #1b1b1b
background container: #131314
Text colour: #E3E3E3


interesting menu svg.

09/08/2025

doing some more colour scheme/css

So done the profile button today. 
Finished the menu solour and the logo div.


Next steps: 

Get div then put profile in it.  green
and put tasks logo. 
and calendar logo in the div. 

add styling to it so hover effect. 


make the tasks the same size, 
add tasks add button. 
make a basic logo, kiss. 
find a logo for all tasks done/no tasks scheduled. 

make a dummy task 
style the dummy task 
scroll bar dummy task test

completed task button 
completed task button styling 
completed task dro down 

add a search bar to completed. etc

10/08/2025

added the profile image, and button green 
Profile button styled done green


11/08/2025

so the navbar inside the page is a broken stratergy, so made it as a component which I should've done from the beginning anyway.

i've made the component and have started to style it, but it is annoying to do this ahh. 



12/08/2025

So things have started speeding up now that i have approached this using everything as a div. which is ugh well it is what it is. 


completed the add task section, completed the completed section at the bottom of the my tasks container. 

Done the sidebar green 



so probably need to do side bar component, green 
so it will move to the right of the screen, green 
and transition should be smooth 

then add divs and sections to it green 
section 1 green 

section 2 all tasks green 
section 3 starred green
section 4 Lists green 
section 5 create new list green  

13/08/2025 

okay so i should really use js and pass alone flags to make the background blue or grey, but I really can't be bothered to do that 
so well actually no duplicating this sort of sidebar 3 or 4 or 5 time is seriously dumb, i should set up the flags now else it's just 
really really dumb to let this problem fester bu dupliacting the code. 


task button navbar when hover, colour change to the rgba 174,203,250,0.5 
make sure the background of the svg of the my tasks all taks near it is solid colour change properly

28/08/2025 
Make the modal stylish for create new list pop up  -- green 
modal done button grey out, 
modal done blue when name entered 

30/08/2025 

going to do the dropdown for the tasks html  -- green 
basic css for the dropdown -- green 
css polished for the dropdown 

highlight for picked mode:  background: rgba(0,74,119,0.8);
light blue: #a8c7fa
background: #004a77; -- background change for highlight picked mode


did some of the task container, and the styling. working on the title for it atm. 


