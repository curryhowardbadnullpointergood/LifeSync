<script>

import './tasks.scss';
import Navbar from '../../lib/components/Navbar.svelte';
import Sidebar from '../../lib/components/Sidebar.svelte';
import { onMount, onDestroy, tick } from 'svelte';
import { goto } from '$app/navigation';


let isSidebarOpen = false

// this is for the tasks make this dynamic in the future 
let open1 = false; 
let open2 = false; 
let open3 = false; 
let open4 = false; 

let isDropdownOpen = false;


// handles the add task temp div 
let showTempTask = false;
// handles the title input for the add new tasks temp div 
let titleInput = '';
// handles the details input for the add new tasks temp div 
let detailsInput = '';
let taskDate = 'null';
let taskTime = 'null';
let showRepeatDropdown = false;
let repeatNumVal = null;
let showRangeDropdown = false;
let taskEndDate = 'null';

let tempAddDiv;
let ignoreNextClick = false;
let repeatVal= 'null';
let repeatOccur = null;
// this is like never, after or on for repeat ends 
let endMode = '';
let repeatNev = null;
let repeatEndsOn ='null';

$: repeatNev = (endMode === 'never') ? true : null;



// this is for getting tasks
// simple tasks just title and details 
let tasksimple = [];
// keeps track of open tasks 
let openTaskId = null;
let openTaskId2 = null;
let openTaskId3 = null;
// this is for displaying todays date dynamically 
const today = new Date();
const tomorrow = new Date(today);
tomorrow.setDate(today.getDate() + 1);

const todayFormatted = formatDate(today);
const tomorrowFormatted = formatDate(tomorrow);
// i guess this is tasks range new now but for today and tomorrow 
let todayTasks = [];
let tomorrowTasks = [];
let todayTasksRepeat = [];
let tomorrowTasksRepeat = [];



// flag for the date modal 
let showDateModal = false;
// flag for the show details in task when button is clicked
let showTaskDetails = false;

// for formatting dates 
function formatDate(d) {
  const yyyy = d.getFullYear();
  const mm = String(d.getMonth() + 1).padStart(2, "0");
  const dd = String(d.getDate()).padStart(2, "0");
  return `${yyyy}-${mm}-${dd}`;
}



  const handleDropdownClick = () => {
    isDropdownOpen = !isDropdownOpen // togle state on click
  }

  const handleDropdownFocusLoss = ({ relatedTarget, currentTarget }) => {
    if (relatedTarget instanceof HTMLElement && currentTarget.contains(relatedTarget)) return  
    isDropdownOpen = false
  }

    async function handleAdd() {
        showTempTask = true;
        ignoreNextClick = true;
    }

    async function loadTasks() {
        const res = await fetch('http://localhost:8090/tasks/simple');
        tasksimple = await res.json();
    }

    async function loadRangeTasks(date, date2) {
         const res = await fetch(
            `http://localhost:8090/tasks/range?date=${date}`
          );
         todayTasks = await res.json();
         
         const res1 = await fetch(
            `http://localhost:8090/tasks/range?date=${date2}`
          );
         tomorrowTasks = await res1.json();
    }

    async function loadRepeatTasks(date, date2) {
        console.log("load repeat tasks function called!")
         const res = await fetch(
            `http://localhost:8090/tasks/repeat?date=${date}`
          );
         todayTasksRepeat = await res.json();
         
         const res1 = await fetch(
            `http://localhost:8090/tasks/repeat?date=${date2}`
          );
         tomorrowTasksRepeat = await res1.json();
    }



    async function completeTask(task, instancedate) {
        // passed 
        console.log("Complete button has been clicked! sanity test ", task)
          try {
            const res = await fetch("http://localhost:8090/tasks/complete", {
              method: "POST",
              headers: {
                "Content-Type": "application/json"
              },
              body: JSON.stringify({
                id: task.id,
                title: task.title,
                details: task.details,
                time: task.time ?? "",
                date: task.date ?? "",
                enddate: task.enddate ?? "",
                instancedate: instancedate
              })
            });
        console.log("response status:", res.status);

        const text = await res.text();
        console.log("response body:", text);

            if (!res.ok) {
              console.error("Failed to complete task");
              return;
            }
        
        
          } catch (err) {
            console.error("Complete task error:", err);
          }
     }
        
        


    async function submitTask() {

          const url =
            'http://localhost:8090/addtask?title='
            + encodeURIComponent(titleInput)
            + '&details=' + encodeURIComponent(detailsInput)
            + '&date=' + encodeURIComponent(taskDate)
            + '&time=' + encodeURIComponent(taskTime)
            + '&repeat=' +encodeURIComponent(repeatVal)
            + '&enddate=' + encodeURIComponent(taskEndDate)
            + '&numrepeat=' + encodeURIComponent(repeatNumVal)
            + '&repeatoccurrences=' + encodeURIComponent(repeatOccur)
            + '&repeatnever=' + encodeURIComponent(repeatNev)
            + '&repeatenddate=' + encodeURIComponent(repeatEndsOn);
        
          console.log("REQUEST URL:", url);
        
          const res = await fetch(url, { method: 'POST' });
        
          console.log("Title:", titleInput);
          console.log("Details:", detailsInput);
          console.log("Date:", taskDate);
          console.log("Time:", taskTime);
          console.log("Repeat: ", repeatVal);
          console.log("End Date:", taskEndDate)
          console.log("Num Repeat: ", repeatNumVal)
          console.log("Repeat Occurrences: ", repeatOccur)
          console.log("Repear Never: ", repeatNev)
          console.log("Repeat End Date: ", repeatEndsOn)
          
          if (res.ok) {
            console.log('Task added!');
            showTempTask = false;
            titleInput = '';
            detailsInput = '';
            taskDate = '';
            taskTime = '';
            repeatVal = '';
            taskEndDate = '';
            repeatNumVal ='';
            repeatOccur = '';
            repeatNev ='';
            repeatEndsOn ='';

          } else {
            console.error('Failed to add task:', await res.text());
          }
        }


    async function finishRangeTask(task, instancedate) {
          const res = await fetch("http://localhost:8090/tasks/completerange", {
          method: "POST",
        headers: { "Content-Type": "application/json" },
              
        body: JSON.stringify({
                id: task.id,
                title: task.title,
                details: task.details,
                time: task.time ?? "",
                date: task.date ?? "",
                enddate: task.enddate ?? "",
                instancedate: instancedate
              })
      });
    
    }
    

    function finalizeTempTask() {
          showTempTask = false;
          submitTask();
    }

        

    function handleClickOutside(event) {
      if (ignoreNextClick) {
        ignoreNextClick = false;
        return;
      }
      if (!showDateModal && showTempTask && tempAddDiv && !tempAddDiv.contains(event.target)) {

        showTempTask = false;
        submitTask();
      }
    }
  onMount(() => {
    if (typeof document !== 'undefined') {
      document.addEventListener('click', handleClickOutside);
    }

    
  // calls tasks get simple tasks 
  loadTasks();
  // calls range tasks gets them based on current time this needs to be changed a bit but not too bad for now 
  loadRangeTasks(todayFormatted, tomorrowFormatted);
  loadRepeatTasks(todayFormatted, tomorrowFormatted);


  });

  onDestroy(() => {
    if (typeof document !== 'undefined') {
      document.removeEventListener('click', handleClickOutside);
    }
  });

  // calls tasks get simple tasks 
 // loadTasks();
  // calls range tasks gets them based on current time this needs to be changed a bit but not too bad for now 
  //loadRangeTasks(todayFormatted, tomorrowFormatted);

</script>



<div class="tasks"> 

  <Navbar on:toggleSidebar={() => isSidebarOpen = !isSidebarOpen}/>
  <Sidebar open={isSidebarOpen} />
  
  <div class="mytasks"> 

    <div class="heading">

      <div class="title">My Tasks</div>

      <div class="dropdown-wrapper" on:focusout={handleDropdownFocusLoss}>
         <button class="dropDownBtn" aria-label="dropDownBtn" on:click={handleDropdownClick} > 
           <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-160q-33 0-56.5-23.5T400-240q0-33 23.5-56.5T480-320q33 0 56.5 23.5T560-240q0 33-23.5 56.5T480-160Zm0-240q-33 0-56.5-23.5T400-480q0-33 23.5-56.5T480-560q33 0 56.5 23.5T560-480q0 33-23.5 56.5T480-400Zm0-240q-33 0-56.5-23.5T400-720q0-33 23.5-56.5T480-800q33 0 56.5 23.5T560-720q0 33-23.5 56.5T480-640Z"/></svg>
        </button>

      {#if isDropdownOpen}
        <div class="dropdown">
          <div class="sortBy"> 
            <div class="heading"> </div>
            <div class="order"> </div>
            <div > </div>
            <div> </div>
            <div> </div>
          </div>
          <div class="listOp"> </div>
          <div class="archiveTasks"> </div>
        </div>
      {/if}

      </div>

    

    </div>

    <div class="addtasks" on:click={handleAdd}> 

      <button aria-label="addtask"> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#a8c7fa"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q65 0 123 19t107 53l-58 59q-38-24-81-37.5T480-800q-133 0-226.5 93.5T160-480q0 133 93.5 226.5T480-160q32 0 62-6t58-17l60 61q-41 20-86 31t-94 11Zm280-80v-120H640v-80h120v-120h80v120h120v80H840v120h-80ZM424-296 254-466l56-56 114 114 400-401 56 56-456 457Z"/></svg>
      </button>

      <p> Add a task</p>

    </div>

    <div class="pendingtasks"> 
        {#if showTempTask}
         <div class="tempaddclass" bind:this={tempAddDiv}>

            <div class="top"> 
              <button> 
<svg xmlns  ="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

              </button>

                  
              <input type="text" placeholder="Title" bind:value={titleInput}>

            </div>

            <div class="middle">
              <input type="text" placeholder="Details" bind:value={detailsInput}>
            </div>

            <div class="bottom"> 
                <button on:click={() => showDateModal = true}>Date</button>
                <button on:click={finalizeTempTask} >Done</button>

            </div> 

         </div> 
        {/if}

        {#each tasksimple as task}
          <div class="taskDisplayed">

            <div class="taskRow">
                <button 
                    class:active={openTaskId3 === task.id}
                    on:click={() =>
                      openTaskId3 === task.id
                        ? (openTaskId3 = null)
                        : (openTaskId3 = task.id)
                    }
                  >
 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
                </button>
        
                <p>{task.title}</p>
            </div>
                
                {#if openTaskId3 === task.id}
                  <div class="taskDetails">
                    <h3>Details</h3>
                    <p>{task.details}</p>
                    <button on:click={() => completeTask(task)}>
                        Completed
                    </button>
                    <button on:click={() => finishRangeTask(task)}>
                        Finish
                    </button>
                    <button on:click={() => goto(`/tasks/${task.id}?instance=${todayFormatted}`)}>
                        Expand
                    </button>
                  </div>
                {/if}
            

          </div>
        {/each}


        {#each todayTasks as task}
        <div class="dateline"> 

            <p>{todayFormatted}</p> 

        </div>

              <div class="taskDisplayed">
            
                <div class="taskRow">
                  <button
                    class:active={openTaskId === task.id}
                    on:click={() =>
                      openTaskId === task.id
                        ? (openTaskId = null)
                        : (openTaskId = task.id)
                    }
                  >
                  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
                
                  </button>
            
                  <p>{task.title}</p>
            
                </div>
            
                {#if openTaskId === task.id}
                  <div class="taskDetails">
                    <h3>Details</h3>
                    <p>{task.details}</p>
                    <p>{task.date} → {task.enddate}</p>
                    <button on:click={() => completeTask(task, todayFormatted)}>
                        Completed
                    </button>
                    <button on:click={() => finishRangeTask(task, tomorrowFormatted)}>
                        Finish
                    </button>
                    <button on:click={() => goto(`/tasks/${task.id}?instance=${todayFormatted}`)}>
                        Expand
                    </button>
                  </div>
                {/if}
            
              </div>
         {/each}
            
        {#each todayTasksRepeat as task}
        <div class="dateline"> 

            <p>{todayFormatted}</p> 

        </div>

              <div class="taskDisplayed">
            
                <div class="taskRow">
                  <button
                    class:active={openTaskId === task.id}
                    on:click={() =>
                      openTaskId === task.id
                        ? (openTaskId = null)
                        : (openTaskId = task.id)
                    }
                  >
                  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
                
                  </button>
            
                  <p>{task.title}</p>
            
                </div>
            
                {#if openTaskId === task.id}
                  <div class="taskDetails">
                    <h3>Details</h3>
                    <p>{task.details}</p>
                    <p>{task.date} → {task.enddate}</p>
                    <button on:click={() => completeTask(task, todayFormatted)}>
                        Completed
                    </button>
                    <button on:click={() => finishRangeTask(task, tomorrowFormatted)}>
                        Finish
                    </button>
                    <button on:click={() => goto(`/tasks/${task.id}?instance=${todayFormatted}`)}>
                        Expand
                    </button>
                  </div>
                {/if}
            
              </div>
         {/each}
            

        {#each tomorrowTasks as task}
        <div class="dateline"> 

            <p>{tomorrowFormatted}</p> 

        </div>
        
              <div class="taskDisplayed">
            
                <div class="taskRow">
                  <button
                    class:active={openTaskId2 === task.id}
                    on:click={() =>
                      openTaskId2 === task.id
                        ? (openTaskId2 = null)
                        : (openTaskId2 = task.id)
                    }
                  >
                  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
                
                  </button>
            
                  <p>{task.title}</p>
            
                </div>
            
                {#if openTaskId2 === task.id}
                  <div class="taskDetails">
                    <h3>Details</h3>
                    <p>{task.details}</p>
                    <p>{task.date} → {task.enddate}</p>
                    <button on:click={() => completeTask(task, tomorrowFormatted)}>
                        Completed
                    </button>
                    <button on:click={() => finishRangeTask(task, tomorrowFormatted)}>
                        Finish
                    </button>
                    <button on:click={() => goto(`/tasks/${task.id}?instance=${tomorrowFormatted}`)}>
                        Expand
                    </button>
                  </div>
                {/if}
            
              </div>
         {/each}
            
        {#each tomorrowTasksRepeat as task}
        <div class="dateline"> 

            <p>{tomorrowFormatted}</p> 

        </div>
        
              <div class="taskDisplayed">
            
                <div class="taskRow">
                  <button
                    class:active={openTaskId2 === task.id}
                    on:click={() =>
                      openTaskId2 === task.id
                        ? (openTaskId2 = null)
                        : (openTaskId2 = task.id)
                    }
                  >
                  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
                
                  </button>
            
                  <p>{task.title}</p>
            
                </div>
            
                {#if openTaskId2 === task.id}
                  <div class="taskDetails">
                    <h3>Details</h3>
                    <p>{task.details}</p>
                    <p>{task.date} → {task.enddate}</p>
                    <button on:click={() => completeTask(task, tomorrowFormatted)}>
                        Completed
                    </button>
                    <button on:click={() => finishRangeTask(task, tomorrowFormatted)}>
                        Finish
                    </button>
                    <button on:click={() => goto(`/tasks/${task.id}?instance=${tomorrowFormatted}`)}>
                        Expand
                    </button>
                  </div>
                {/if}
            
              </div>
         {/each}
            


    </div>

    <div class="completedtasks">
      
      <button aria-label="arrowright"> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M400-280v-400l200 200-200 200Z"/></svg>

      </button>

      <p> Completed</p>
    </div>

  </div>


</div>



{#if showDateModal}
    <div class="modal">
        <div class="modalcontent">
            <p>Date</p>
            <input type="date" class="textbox" placeholder="Date" bind:value={taskDate}>
            <p>Set Time</p>
            <input type="time" class="time-input" bind:value={taskTime}>
            <div class="range">
                <button
                    on:click={() => showRangeDropdown = !showRangeDropdown}
                    class:active={showRangeDropdown}
                > 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

                </button>
                <p>Range</p>
            </div>

                {#if showRangeDropdown}
                    <div class="rangeDropdown">
                        <div class="setrange">
                            <p>End Date</p>
                            <input type="date" class="textbox" bind:value={taskEndDate}>
                        </div>
                    </div>
                {/if}

            
            <div class="repeat">
                <button
                    on:click={() => showRepeatDropdown = !showRepeatDropdown}
                    class:active={showRepeatDropdown}
                > 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

                </button>
                <p>Repeat</p>
            </div>

            {#if showRepeatDropdown}
                <div class="repeatDropdown">
                    <div class="top"> 
                        <input type="number" min="1" bind:value={repeatNumVal}>
                          <select name="timeframe" id="timeframe" bind:value={repeatVal}>
                            <option value="day">day</option>
                            <option value="week">week</option>
                            <option value="month">month</option>
                            <option value="year">year</option>
                          </select>
                    </div>

                    <div class="ends"> 
                        <p>Ends</p>

                        <div class="never">
                
                            <button on:click={() => endMode = 'never'} class:active={endMode === 'never'} > 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

                            </button>

                            <p>Never</p>

    
                        </div>
                        <div class="on">
                
                            <button on:click={() => endMode ='on'} class:active={endMode === 'on'}> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

                            </button>

                            <p>On</p>

                            <input type="date" class="textbox" placeholder="Date" bind:value={repeatEndsOn}>

    
                        </div>
                        <div class="after">
                
                            <button on:click={() => endMode ='after'} class:active={endMode === 'after'}> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

                            </button>

                            <p>After</p>

                            <div class="occurrences">
                                <input type="number" min="1" bind:value={repeatOccur}>
                                <p>Occurrences</p>
                            </div>
    
                        </div>

                    </div>
                </div>
            {/if}


            <div class="bottom"> 
                <button class="cancel" on:click={() => showDateModal = false}> Cancel </button>
                <!-- This probably needs to be more nuanced in the future this on clikc for done should save values and set to false-->
                <button class="done" on:click={() => showDateModal = false}> Done </button>
            </div>
                  
                  
        </div>
    </div>
{/if}


