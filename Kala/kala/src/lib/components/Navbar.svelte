<script>

import './navbar.scss'
import { createEventDispatcher } from 'svelte';
import { page } from '$app/stores';


$: isTasksPage = $page.url.pathname === '/tasks';

const dispatch = createEventDispatcher();

function toggle() {
  dispatch('toggleSidebar');
}


</script> 






<header>
  <nav class="navbar">
    <div class="leftSide">
      
      <button aria-label="Menu" on:click={toggle}> 
  
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#E3E3E3"><path d="M120-240v-80h720v80H120Zm0-200v-80h720v80H120Zm0-200v-80h720v80H120Z"/></svg>
  
      </button>

    </div>
    <div class="rightSide">

      <div class="calendar">

        <button aria-label="calendar"> 

          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M200-80q-33 0-56.5-23.5T120-160v-560q0-33 23.5-56.5T200-800h40v-80h80v80h320v-80h80v80h40q33 0 56.5 23.5T840-720v560q0 33-23.5 56.5T760-80H200Zm0-80h560v-400H200v400Zm0-480h560v-80H200v80Zm0 0v-80 80Zm280 240q-17 0-28.5-11.5T440-440q0-17 11.5-28.5T480-480q17 0 28.5 11.5T520-440q0 17-11.5 28.5T480-400Zm-160 0q-17 0-28.5-11.5T280-440q0-17 11.5-28.5T320-480q17 0 28.5 11.5T360-440q0 17-11.5 28.5T320-400Zm320 0q-17 0-28.5-11.5T600-440q0-17 11.5-28.5T640-480q17 0 28.5 11.5T680-440q0 17-11.5 28.5T640-400ZM480-240q-17 0-28.5-11.5T440-280q0-17 11.5-28.5T480-320q17 0 28.5 11.5T520-280q0 17-11.5 28.5T480-240Zm-160 0q-17 0-28.5-11.5T280-280q0-17 11.5-28.5T320-320q17 0 28.5 11.5T360-280q0 17-11.5 28.5T320-240Zm320 0q-17 0-28.5-11.5T600-280q0-17 11.5-28.5T640-320q17 0 28.5 11.5T680-280q0 17-11.5 28.5T640-240Z"/></svg>

        </button>

      </div>

      <div class="task" class:active={isTasksPage}>

        <button aria-label="tasks" class:active={isTasksPage}>  
          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e8eaed"><path class="LuCW2c" d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q65 0 123 19t107 53l-58 59q-38-24-81-37.5T480-800q-133 0-226.5 93.5T160-480q0 133 93.5 226.5T480-160q133 0 226.5-93.5T800-480q0-18-2-36t-6-35l65-65q11 32 17 66t6 70q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm-56-216L254-466l56-56 114 114 400-401 56 56-456 457Z"></path></svg>
        </button>

      </div>


        
    

      <button class="profile">
        <img src="/Ash.jpg" alt="P">
      </button>

    </div>
  </nav>
</header>


<style lang="scss">
  
  header{
    display: flex;
    flex-direction: row;
    
    align-items: center;
    min-height: 4rem;
    min-width: 100%;
    
    // margin: 1rem;
    

    .navbar{
    
      display: flex;
      align-items: center;
      max-height: 3rem;
      width: 100%;
      
      margin-left: 1rem;
      margin-right: 1rem;
      justify-content: space-between;

      .leftSide{
        display: flex;
        flex-direction: row;
        align-items: center;

        button{

         background: none;
         cursor: pointer; 
  
         
  
         min-width: 3rem; 
         min-height: 3rem;
  
         border-radius: 50%;
         border: 2px solid transparent;
  
         // padding: 0px 10px;
       
          
        }
        
        button:hover{
          background-color: #363737;
  
        }


      }

      .rightSide{
        display: flex;
        flex-direction: row;
        align-items: center;

        min-height: 3rem;

        gap: 0.5rem;


        .calendar{
          display: flex;
          align-self: center;

          min-height: 3rem;
          min-width: 3rem;
          button{
            display: flex;
            align-items: center;
            justify-content: center;
            
            background: #1b1b1b;
            border-radius: 50%;
            border: 2px solid transparent;  
            
            min-width: 3rem;


          }
          
          button:hover{
            background: #373636;
          }

        }
        .task{
          display: flex;
          align-self: center;
          justify-content: center;
          background: #1b1b1b;
          
          min-height: 3rem;
          min-width: 3rem;

          button{
            display: flex;
            align-items:  center;
            justify-content: center;

            // background: #004a77;
            background: #1b1b1b;
            border-radius: 50%;
            border: 2px solid transparent;  
            
            min-width: 3rem;
          }

          button:hover{
            background: #373636;

          }
          
          button.active{
            background: rgba(0,74,119,0.8);
          }

          button.active:hover{
            background: #004a77;

          }

        }

        .profile{
          display: flex;
          align-items: center;
          justify-content: center;
          border-radius: 50%;
          border: 0.3rem solid transparent;
          overflow: hidden;
          background: #1b1b1b;

          min-width: 3rem;
          min-height: 3rem;
          
          

          img{
            width: 2.5rem;
            height: 2.5rem;
            object-fit: cover;

            display: block;
  
          }
        }
        
        .profile:hover{
          background: #363737;    
        }

      }
  
    }

  }

</style>


<!-- 

  <div class="navBar"> 

    <div class="menuLogo"> 

      <button aria-label="Menu"> 
  
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#E3E3E3"><path d="M120-240v-80h720v80H120Zm0-200v-80h720v80H120Zm0-200v-80h720v80H120Z"/></svg>
  
      </button>

      <div class="logo">
        LOGO
      </div>
  

    </div> 

    
    <div class="rightSide">

      <div class="todo"> 
        
        <button aria-label="Task">

          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q65 0 123 19t107 53l-58 59q-38-24-81-37.5T480-800q-133 0-226.5 93.5T160-480q0 133 93.5 226.5T480-160q133 0 226.5-93.5T800-480q0-18-2-36t-6-35l65-65q11 32 17 66t6 70q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm-56-216L254-466l56-56 114 114 400-401 56 56-456 457Z"/></svg>


        </button> 
      </div>

      <div class="profile">
  
        <button>
          <img src="/Ash.jpg" alt="P">
        </button>
      </div>
  
    </div>




  </div> 

-->
