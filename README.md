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

Similar to add tasks, but for people. So you can add new people you meet on here, their social media information, contact information, and a note about them, their birthdays, important days etc, Say they tell you they are leaving the country at X time etc, you can remember to say goodbye or smth, or buy a gift etc. This is again sort of like a task, but related to people instead. There's also a note section where you can take note about what they like sort of their favourite foods/colour etc. Will make you seem socially adept lmao. 


Also it can remind you to contact people every now and then so you stay in touch. Perhaps automatic messages like Happy Birthday or Happy Christmas etc. This saves time, and helps you keep track of things. You can then manually respond back to them if they say anything back. 

So this requires some API stuff I assume to make it automatically run in the background. 

Also reminds you to buy gifts for them, so this is based on whether you want to give them a gift or not, and if so it reminds you 2 weeks before, this should be a variable option, obviously more important people you should think about what to give etc. 


----------------
