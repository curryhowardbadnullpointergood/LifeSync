<script>

  import { page } from '$app/stores';
  $: isTasksPage = $page.url.pathname === '/tasks';
  $: isSatrredPage = $page.url.pathname === '/starred';

	export let open = false
</script>

<aside class="sidebar" class:open>
	<nav class="sidNav">
    <div class="createTask"> 
      <div class="add"> 
        
        <button aria-label="addicon">

          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z"/></svg>
        
        </button> 
        
        <p>Create</p>
      </div>
    </div>

    <div class="alltask" class:active={isTasksPage}> 

      <a href="/tasks">All tasks</a>

    </div> 

    <div class="starred" class:active={isSatrredPage}> 
      <a href="/starred"> Starred</a>

    </div> 

    <div class="lists"> 


      <p>Lists</p>
      <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-360 280-560h400L480-360Z"/></svg>

    </div> 

    <div class="newlist"> 

      <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z"/></svg>
      
      <p>Create new list</p>

    </div> 

	</nav>
</aside>

<style lang="scss">

  .sidebar{
    position: fixed;
    background: #1b1b1b;
    left: -100%;

    margin-top: 5rem;
    min-height: 100vh;
    min-width: 16rem;

    display: flex;
    flex-direction: column;

    
    .sidNav{
      display: flex;
      flex-direction: column;

      
      margin-left: 1rem;
      gap: 0.5rem;
      


      .createTask{
        display: flex;
        flex-direction: row;
        align-items: center;
        
        min-height: 3.5rem;
        background: #1b1b1b;
      
        

        .add{
          display: flex;
          flex-direction: row;
          align-items: center;
          //justify-content: center;
          

          min-height: 3.5rem;
          width: 7.25rem;
          //background: #37393b;
          background: rgba(54, 55, 55, 0.8);

          border-radius: 1rem;
          overflow: hidden;

          button{
            display: flex;
            flex-direction: row;
            background: transparent;
            padding: 0.5rem;
            
          }
          

  
          p{
            padding-top: 4px;
            color: #e3e3e3;
            font-family: "Lucida Console", "Courier New", monospace;
            font-weight: 400;
            
          }
        }

        .add:hover{
          background: #37393b;


        }
        
      }


      .alltask{
        margin-top: 1rem;
        padding-left: 1rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        min-height: 2rem;
        background: #1b1b1b;
        border-radius: 1.5rem;

        a{
          text-decoration: none;
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 0.875rem;
        }
      }

      .alltask:hover{
        background: #363737;
      }
      .alltask.active:hover{
        background: #004a77;
      }
      .alltask.active{
        background: rgba(0,74,119,0.8);

      }
      
      .starred{
        padding-left: 1rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        min-height: 2rem;
        border-radius: 1.5rem;
        background: #1b1b1b;
        a{
          text-decoration: none;
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 0.875rem;
        }
      }
      .starred:hover{
        background: #363737;
      }

      .starred.active:hover{
        background: #004a77;
      }
      .starred.active{
        background: rgba(0,74,119,0.8);

      }

      .lists{
        padding-left: 1rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        min-height: 2rem;
        border-radius: 1.5rem;
        background: #1b1b1b;

        gap: 6rem;
        p{
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 0.875rem;
        }
        
        svg{
          border-radius: 50%;
          border: solid 0.5rem transparent;
        }
        svg:hover{
          background: #363737;
        }
      }
      .newlist{
        padding-left: 1rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        min-height: 2rem;
        border-radius: 1.5rem;
        background: #1b1b1b;

        gap: 0.5rem; 
        p{
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 0.875rem;
        }
      }
      .newlist:hover{
        background: #363737;
      }

    }

    
  }

  .sidebar.open {
    left: 0;
    // transform: translateX(0);
  }
	
</style>
