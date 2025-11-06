
<script>

//import './tasks.scss';
//import Navbar from '../../lib/components/Navbar.svelte';
//import Sidebar from '../../lib/components/Sidebar.svelte';
import { onMount, onDestroy } from 'svelte';



let isSidebarOpen = false

// this is for the tasks make this dynamic in the future 
let open1 = false; 
let open2 = false; 
let open3 = false; 
let open4 = false; 

let isDropdownOpen = false;


  const handleDropdownClick = () => {
    isDropdownOpen = !isDropdownOpen // togle state on click
  }

  const handleDropdownFocusLoss = ({ relatedTarget, currentTarget }) => {
    if (relatedTarget instanceof HTMLElement && currentTarget.contains(relatedTarget)) return  
    isDropdownOpen = false
  }


</script>



<div class="tasks"> 

  
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

    <div class="addtasks"> 

      <button aria-label="addtask"> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#a8c7fa"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q65 0 123 19t107 53l-58 59q-38-24-81-37.5T480-800q-133 0-226.5 93.5T160-480q0 133 93.5 226.5T480-160q32 0 62-6t58-17l60 61q-41 20-86 31t-94 11Zm280-80v-120H640v-80h120v-120h80v120h120v80H840v120h-80ZM424-296 254-466l56-56 114 114 400-401 56 56-456 457Z"/></svg>
      </button>

      <p> Add a task</p>

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


