<script>

import './tasks.scss'


// this is for the tasks make this dynamic in the future 
let open1 = false; 
let open2 = false; 
let open3 = false; 
let open4 = false; 


</script>



<div class="tasks"> 

  <div class="navBar"> 

    <div class="menuLogo"> 

      <button aria-label="Menu"> 
  
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#E3E3E3"><path d="M120-240v-80h720v80H120Zm0-200v-80h720v80H120Zm0-200v-80h720v80H120Z"/></svg>
  
      </button>

      <div class="logo">
        LOGO
      </div>
  

    </div> 

    

          
    <button class="profileButton" aria-label="Open user menu">
    </button>




  </div> 

  <div class="taskList">

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
    

</div>

