<script lang="ts">
    import Header from "$lib/Header.svelte";
    import PostCard from "$lib/Post.card.svelte";
    import PostPopularWidget from "$lib/Popular.posts.svelte";
  
    import { appService } from "$lib/DataService";
    import type { AppUser } from "$lib/DataInterface";
 
    let localUser: AppUser = undefined;
    appService.user.subscribe((value) => {
      localUser = value;
    });

    let handleInput: string = "";
    let handleExists: boolean = false;
    
    function testIfHandleExists() {
      appService.GetIfHandleExists(handleInput).then((exists) => {
        handleExists = exists;
      });
    }

    function cancel() {
      submitHandle(localUser.handle);
    }

    function submitHandle(newHandle: string) {
      appService.SetHandle(newHandle).then((result) => {
        appService.Navigate("/home")
      });
    }
  </script>
  
  <div>
    <Header />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="popup_dialog" on:click={cancel}>
      <div class="popup_dialog_content" on:click|stopPropagation={() => {}}>
          <h1>Welcome {localUser.handle}!</h1>
          <br/>
          <div>
            Your handle was automatically generated. You can set a new handle that is available, if you wish.
            <br /><br />
            <input placeholder="New user handle" autofocus class:handleExists
              bind:value={handleInput}
              on:keyup|stopPropagation={testIfHandleExists}
            />
            <br />
            {#if handleInput && handleExists}
              <div class="handle_tip_negative">
                Handle already taken!
              </div>
            {:else if handleInput}
              <div class="handle_tip_positive">
                Handle available!
              </div>            
            {/if}
            <br />
            <button class="submit_button button_big_primary" on:click={() => submitHandle(handleInput)}>Save</button>
          </div>
      </div>
    </div>
  </div>
  
  <style>
    .popup_dialog {
        position: fixed;
        height: 100vh;
        width: 100vw;
        top: 0px;
        left: 0px;
        background-color: #6c6c6cb3;
        z-index: 100;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .popup_dialog_content {
        position: absolute;
        /* top: 40px;
        bottom: 40px; */
        height: 204px;
        width: 90%;
        max-width: 464px;
        background: white;
        margin: auto;
        border: 1px #eee solid;
        border-radius: 24px;
        font-size: 14px;
        padding: 28px;
        padding-top: 12px;
        padding-right: 54px;
        overflow-y: auto;
    }

    .popup_dialog_content h1 {
      font-size: 24px;
    }

    .handleExists {
      border-color: rgb(201, 74, 74);
    }

    .handle_tip_positive {
      margin-top: 4px;
      margin-left: 2px;
      color: rgb(26, 137, 23);
      font-size: 12px;
    }

    .handle_tip_negative {
      margin-top: 4px;
      margin-left: 2px;
      color: rgb(201, 74, 74);
      font-size: 12px;
    }

    .submit_button {
      right: 44px;
      position: absolute;
      bottom: 20px;
    }

  </style>
  