<script>
  import { onMount } from "svelte";
  import nbaTeams from "../types/nbaTeams.js";
  import { fetchFantasyTeamRosters } from "../api/fantasy_team_rosters.js";

  let fantasyTeamRosterData = $state(null);
  let loading = $state(true);
  let error = $state(null);

  onMount(async () => {
    try {
      loading = true;
      error = null;

      const fantasyTeamRosterResult = await fetchFantasyTeamRosters();

      fantasyTeamRosterData = fantasyTeamRosterResult;
    } catch (err) {
      console.error("Error loading fantasy team roster data:", err);
      error = err.message;
    } finally {
      loading = false;
    }
  });
</script>

{JSON.stringify(fantasyTeamRosterData, null, 2)}

<style></style>
