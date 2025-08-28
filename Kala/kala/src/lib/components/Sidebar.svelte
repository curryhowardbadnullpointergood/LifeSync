<script>

  import { page } from '$app/stores';
  $: isTasksPage = $page.url.pathname === '/tasks';
  $: isSatrredPage = $page.url.pathname === '/starred';

	export let open = false

  let showModal = false;

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
      <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="m354-287 126-76 126 77-33-144 111-96-146-13-58-136-58 135-146 13 111 97-33 143ZM233-120l65-281L80-590l288-25 112-265 112 265 288 25-218 189 65 281-247-149-247 149Zm247-350Z"/></svg>
      
      <a href="/starred"> Starred</a>

    </div> 

    <div class="lists"> 


      <p>Lists</p>
      <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-360 280-560h400L480-360Z"/></svg>

    </div> 

    <div class="newlist"> 
      
      <button on:click={() => showModal = true}> 
        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z"/></svg>
      


      <p>Create new list</p>

      </button>



	</nav>
</aside>

{#if showModal}
  <div class="modal-overlay">
    <div class="modal">

      <div class="top">

        <p>Create new list</p>

      </div>

      <div class="middle">
        <input type="text" class="textbox" placeholder="Enter name" required>
      </div>

      <div class="holder">

      </div>
      <div class="bottom">
        <button class="cancel" on:click={() => showModal = false}> Cancel </button>
        <button class="done"> Done </button>
      </div>
    </div>
  </div>
{/if}



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
          //padding-top: 0.5rem;
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
        gap: 0.5rem;
        a{
          padding-top: 0.4rem;
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

        

        button{
          background: none;
          display: flex;
          flex-direction: row;
          gap: 0.5rem;
         p{
           padding-top: 0.3rem;
           color: #e3e3e3;
           font-family: "Lucida Console", "Courier New", monospace;
           font-weight: 400;
           font-size: 0.875rem;
         }
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
	
  .modal-overlay {
    display: flex; 
    flex-direction: column;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5); /* semi-transparent black */
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 999; /* High z-index to ensure it overlays everything */

    .modal {
      display: flex; 
      flex-direction: column;
      background: #1e1e1e;
      border-radius: 8px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
      color: white;
      
      max-height: 9rem;
      min-height: 8.99rem;
      min-width: 17rem;
      max-width: 17.001rem;

      padding: 1rem;

      align-items: center;

      .top{
        display: flex;
        flex-direction: row;
        align-items: left;
        //background: red;
        width: 100%;

        p{
          padding-top: 4px;
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;

        }
      }


      .bottom{
        background: #1e1e1e;
        
        padding: 1rem;
        
        .cancel{
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 16px;
      
          background: #1e1e1e;
          border: 1px solid transparent;
          border-radius: 10px;
          color: #a8c7fa;
          cursor: pointer;
          padding: 5px 10px;
          
        }
      
        .cancel:hover{
          background: rgba(174,203,250,0.3);
        }

        
        
        
        .done{
          background: #1e1e1e; 
          color: #e3e3e3;
          font-family: "Lucida Console", "Courier New", monospace;
          font-weight: 400;
          font-size: 16px;
      
          //background: red;
          border: 1px solid transparent;
          border-radius: 10px;
          color: #a8c7fa;
          cursor: pointer;
          padding: 5px 10px;
          

        }

        
      
      }

      .middle{
        //background: red;
        padding: 1rem;
        width: 100%;
        .textbox{
          border: none; 
          background: none; 
          width: 100%;
          border-bottom: 2px solid #a8c7fa;
          padding-top: 1.5rem;
          border-radius: 3px;
          color: #e3e3e3;
          outline: none;
        }

        .textbox::placeholder {
          color: #e3e3e3;
        }
      }
      
      .holder{
        //min-height: 1rem;
        min-width: 20px;
        //background: red;
      }


    }
  }
</style>
