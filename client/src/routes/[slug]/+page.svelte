<script lang="ts">
  import Header from "../../lib/Header.svelte";
  import PostCard from "../../lib/Post.card.svelte";
  import PostPopularWidget from "../../lib/Popular.posts.svelte";

  import { appService } from "$lib/DataService";
  import type {
    PostOverview,
    PostOverviewCollection,
  } from "$lib/DataInterface";
  import { ToTitleCase } from "$lib/DataInterface";

  export let data: { userId: string; posts: PostOverview[] };

  let start: number = 0;
  let limit: number = 5;

  function scrollCheckEnd(event) {
    console.log(event);

    if (
      event.target.scrollTop >=
      event.target.scrollHeight - event.target.clientHeight
    ) {
      console.log("scroll end");

      start = start + limit;
      appService.GetUserPosts(data.userId, start, limit).then((result) => {
        data.posts = data.posts.concat(result);
      });
    }
  }
</script>

<div class="page_box" on:scroll={scrollCheckEnd}>
  <Header />

  <div class="content">
    {#if data && data.posts}
      <div class="container">
        <div class="panel_left">
          <div class="panel_left_inner">
            <div class="tag_title">@{data.userId}</div>
            <div class="user_header">
              <table class="styled_table">
                <tbody>
                  <tr>
                    <td>Joined</td>
                    <td style="text-align: right;">Aug 23rd, 2023</td>
                  </tr>
                  <tr>
                    <td>Posts</td>
                    <td style="text-align: right;">12</td>
                  </tr>
                  <tr>
                    <td>Followers</td>
                    <td style="text-align: right;">6000</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <h1>Posts</h1>
            {#each data.posts as post}
              <div>
                <PostCard {post} />
              </div>
            {/each}
            <div class="pannel_left_footer" />
          </div>
        </div>
      </div>
    {/if}
  </div>
</div>

<style>
  .page_box {
    height: calc(100vh);
    overflow-y: auto;
  }

  .content {
    max-width: 1336px;
    text-align: left;
    margin: auto;
    /* height: 100%;
    height: calc(100vh - 58px);
    overflow-y: hidden; */
  }

  .container {
    display: flex;
    justify-content: space-evenly;
    flex-direction: row;
    /*
    height: calc(100vh - 58px);
    overflow-y: auto; */
  }

  .panel_left {
    width: 100%;
    flex: 1 1 auto;
    justify-content: center;
    display: inline-flex;
    /* height: 100%;
    overflow-y: auto; */
  }

  .tag_title {
    margin-top: 30px;
    margin-left: 24px;
    font-size: 48px;
    font-weight: 600;
  }

  .panel_left_inner {
    max-width: 728px;
    width: 100%;
    padding-bottom: 54px;
  }

  .pannel_left_footer {
    height: 84px;
  }

  .panel_right {
    /* min-height: 100vh; */

    border-left: 1px solid rgba(242, 242, 242, 1);
    padding-left: 32px;
    min-width: 420px;

    /* uncomment below to make right panel sticky
    height: calc(100vh - 58px);
    position: sticky;
    top: 58px;
    overflow-y: auto; */
  }

  .user_header {
    padding: 24px;
  }

  .widget1 {
    height: 100%;
    width: 100%;
  }

  h1 {
    margin-left: 22px;
    margin-bottom: 0px;
  }

  .styled_table {
    width: 100%;
    max-width: 444px;
  }

  @media (max-width: 903.98px) {
    /* .cc {
      min-width: 0;
    }

    .sb {
      min-width: 0;
    } */

    .panel_right {
      display: none;
    }
  }

  @media (min-width: 904px) and (max-width: 1079.98px) {
    /* .cc {
      max-width: 680px;
    }

    .sb {
      max-width: 352px;
      min-width: 310px;
    } */
  }

  @media (min-width: 1080px) {
    /* .sb {
      max-width: 352px;
      min-width: 352px;
      padding-right: 24px;
    } */
  }

  @media (min-width: 1080px) {
    /* .sb {
      padding-left: clamp(24px, 24px + 100vw - 1080px, 40px);
    } */
  }
</style>
