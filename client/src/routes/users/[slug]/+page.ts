import { error } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import { appService } from "../../../lib/DataService";
import { UserState } from "$lib/DataInterface";

export const load: PageLoad = (({ params }) => {
  return {
    userId: params.slug,
    posts: appService.GetUserPosts(params.slug, 0, 5)
  };
}) satisfies PageLoad;
