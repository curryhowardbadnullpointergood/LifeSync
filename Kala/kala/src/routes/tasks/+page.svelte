<script>

import './tasks.scss';
import Navbar from '../../lib/components/Navbar.svelte';
import Sidebar from '../../lib/components/Sidebar.svelte';
import { onMount, onDestroy, tick } from 'svelte';


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
let taskDate = '';
let taskTime = '';
let showRepeatDropdown = false;
let repeatNumVal = 1;
let showRangeDropdown = false;
let taskEndDate = '';

let tempAddDiv;
let ignoreNextClick = false;
let repeatVal= '';
let repeatOccur=1;
// this is like never, after or on for repeat ends 
let endMode = 'never';



// flag for the date modal 
let showDateModal = false;

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
            + '&repeatoccurrences=' + encodeURIComponent();
        
          console.log("REQUEST URL:", url);
        
          const res = await fetch(url, { method: 'POST' });
        
          console.log("Title:", titleInput);
          console.log("Details:", detailsInput);
          console.log("Date:", taskDate);
          console.log("Time:", taskTime);
          console.log("Repeat: ", repeatVal);
          console.log("End Date:", taskEndDate)
          console.log("Num Repeat: ", repeatNumVal)
          console.log("Repeat Occurrences: ", repeatoccurrences)
          
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

          } else {
            console.error('Failed to add task:', await res.text());
          }
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
  });

  onDestroy(() => {
    if (typeof document !== 'undefined') {
      document.removeEventListener('click', handleClickOutside);
    }
  });

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
            </div> 

         </div> 
        {/if}

        <div class="taskDisplayed">

            <button> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

            </button>


           <p> Example text here: <p>
        </div>
        <div class="taskDisplayed">

            <button> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>

            </button>


           <p> Example text here: <p>
        </div>


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
            <input type="date" class="textbox" placeholder="Date" required bind:value={taskDate}>
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

                            <input type="date" >

    
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


