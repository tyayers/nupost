<script lang="ts">
  import type { AppUser } from "./DataInterface";
  import { appService } from "./DataService";

  export let onEnter: () => void;
  export let onLeave: () => void;
  export let followingChanged: (following: boolean) => void;

  export let user: AppUser
  export let localUser: AppUser
  export let userFollowing: boolean
  export let size: string

  function onMouseEnter() {
    if (onEnter) onEnter();
  }

  function onMouseLeave() {
    if (onLeave) onLeave();
  }

  function follow() {
    appService.FollowUser(user.uid);
    userFollowing = true;
    followingChanged(userFollowing);
  }

  function unFollow() {
    appService.UnFollowUser(user.uid);
    userFollowing = false;
    followingChanged(userFollowing);
  }
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class={"popup_menu_panel_" + size }
    on:mouseenter={onMouseEnter}
    on:mouseleave={onMouseLeave}>
  <div class={"popup_menu_arrow " + "popup_menu_arrow_position_" + size} />
  <div class="popup_menu">
    <div class="popup_menu_inner">
      {user.displayName}

      {#if user.uid != localUser.uid && !userFollowing}
        <button on:click|preventDefault={follow}>Follow</button>
      {:else if user.uid != localUser.uid && userFollowing}
        <button on:click|preventDefault={unFollow}>Unfollow</button>
      {/if}

    </div>
  </div>
</div>

<style>
  .popup_menu_panel_small {
    position: absolute;
    margin-top: 34px;
    z-index: 2;
  }

  .popup_menu_panel_large {
    position: absolute;
    margin-top: 48px;
    z-index: 2;
  }

  .popup_menu {
    position: relative;
    left: -14px;
    top: -10px;
    height: 200px;
    width: 260px;
    overflow-y: auto;
    border-radius: 3px;
    background: rgb(255, 255, 255);
    box-shadow: rgba(0, 0, 0, 0.15) 0px 2px 10px 0px;
    border: 1px solid rgb(242, 242, 242);
    border-radius: 4px;
  }

  .popup_menu_inner {
    position: relative;
    background: rgb(255, 255, 255);
    width: 100%;
    z-index: 2;
    padding-top: 20px;
    /* padding-bottom: 20px; */
    font-family: sohne, "Helvetica Neue", Helvetica, Arial, sans-serif;
  }

  .popup_menu_arrow {
    position: relative;
    top: -22px;
    
    z-index: 1;
    border: 1px solid rgb(242, 242, 242);
    box-shadow: rgba(0, 0, 0, 0.15) -1px -1px 1px -1px;
    transform: rotate(45deg) translate(16px, 16px);
    background: rgb(255, 255, 255);
    height: 14px;
    width: 14px;
    display: block;
    content: "";
    pointer-events: none;
  }

  .popup_menu_arrow_position_small {
    left: 4px;
  }

  .popup_menu_arrow_position_large {
    left: 14px;
  }

  .popup_menu_item {
    padding-top: 5px;
    padding-bottom: 16px;
    padding-left: 10px;
    /* border-bottom: 1px dashed rgb(242, 242, 242); */
    cursor: pointer;
    font-size: 16px;
  }
</style>
