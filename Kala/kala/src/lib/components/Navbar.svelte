<script>

import './navbar.scss'

</script> 






<header>
  <nav class="navbar">
    <div class="leftSide">
      
      <button aria-label="Menu"> 
  
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#E3E3E3"><path d="M120-240v-80h720v80H120Zm0-200v-80h720v80H120Zm0-200v-80h720v80H120Z"/></svg>
  
      </button>

    </div>
    <div class="rightSide">
      LOGO
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
