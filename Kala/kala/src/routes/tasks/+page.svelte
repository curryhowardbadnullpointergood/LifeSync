<script>

import './tasks.scss';
import Navbar from '../../lib/components/Navbar.svelte';

// this is for the tasks make this dynamic in the future 
let open1 = false; 
let open2 = false; 
let open3 = false; 
let open4 = false; 


</script>



<div class="tasks"> 

  <Navbar />



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


