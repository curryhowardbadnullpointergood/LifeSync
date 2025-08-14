<script>

import './starred.scss';
import Navbar from '../../lib/components/Navbar.svelte';
import Sidebar from '../../lib/components/Sidebar.svelte';

let isSidebarOpen = false

// this is for the tasks make this dynamic in the future 
let open1 = false; 
let open2 = false; 
let open3 = false; 
let open4 = false; 


</script>



<div class="tasks"> 

  <Navbar on:toggleSidebar={() => isSidebarOpen = !isSidebarOpen}/>
  <Sidebar open={isSidebarOpen} />
  
  <div class="mytasks"> 

    <div class="heading">

      <div class="title">Starred Tasks</div>

    </div>

    <div class="addtasks"> 

      <button aria-label="addtask"> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#a8c7fa"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q65 0 123 19t107 53l-58 59q-38-24-81-37.5T480-800q-133 0-226.5 93.5T160-480q0 133 93.5 226.5T480-160q32 0 62-6t58-17l60 61q-41 20-86 31t-94 11Zm280-80v-120H640v-80h120v-120h80v120h120v80H840v120h-80ZM424-296 254-466l56-56 114 114 400-401 56 56-456 457Z"/></svg>
      </button>

      <p> Add a starred task</p>

    </div>

    <div class="pendingtasks"> 

    </div>

    <div class="completedtasks">
      
      <button aria-label="arrowright"> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M400-280v-400l200 200-200 200Z"/></svg>

      </button>

      <p> Completed</p>
    </div>

  </div>


</div>

  <!-- <div class="taskList">

    <div class="individualTask">

      
      <div class="taskOverview" on:click={() => open1 = !open1}>
        <h4> Task 1: </h4>
        <p> Points: X </p>
        <i> EST Time: 10 hours </i> 
        <b> Pirority: High</b>
      </div>


      <div class="taskDescriptionShort">
        <b> Task overview, high level overview, for more details have to click! </b>
      </div> 
    
      {#if open1}
        <div class="taskContent">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
        </div>
      {/if}

    </div>

    <div class="individualTask">
    
          
          <div class="taskOverview" on:click={() => open2 = !open2}>
            <h4> Task 2: </h4>
            <p> Points: X </p>
            <i> EST Time: 10 hours </i> 
            <b> Pirority: High</b>
          </div>
    
    
          <div class="taskDescriptionShort">
            <b> Task overview, high level overview, for more details have to click! </b>
          </div>

       {#if open2}
        <div class="taskContent">
          <p>This is the detailed content of Description 2.</p>
        </div>
      {/if}
    
      </div>


     <div class="individualTask">
    
          
          <div class="taskOverview" on:click={() => open3 = !open3}>
            <h4> Task 3: </h4>
            <p> Points: X </p>
            <i> EST Time: 10 hours </i> 
            <b> Pirority: High</b>
          </div>
    
    
          <div class="taskDescriptionShort">
            <b> Task overview, high level overview, for more details have to click! </b>
      </div> 

          {#if open3}
              <div class="taskContent">
                <p>This is the detailed content of Description 2.</p>
              </div>
          {/if} 
      </div>


     <div class="individualTask">
    
          
          <div class="taskOverview" on:click={() => open4 = !open4}>
            <h4> Task 4: </h4>
            <p> Points: X </p>
            <i> EST Time: 10 hours </i> 
            <b> Pirority: High</b>
          </div>
    
    
          <div class="taskDescriptionShort">
            <b> Task overview, high level overview, for more details have to click! </b>
          </div> 
          
           {#if open4}
            <div class="taskContent">
              <p>This is the detailed content of Description 2.</p>
            </div>
          {/if}
        
        </div>
     
        
      </div>
    
  -->



