<script lang="ts">
  import { page } from '$app/stores';
	import * as Card from '$lib/components/ui/card';

  export let name: string = "im pasta";
  export let description: string = "you are not pasta";
  export let picture: string = "https://i.imgur.com/T5zyE63.png";
  export let address: string = "pun street";
</script>

<a href="{$page.route.id}/{name}">
  <Card.Root class="w-[399px] h-[385px]">
    <Card.Content>
      <img src={picture} alt="" class="rounded-[11px] w-[399px] h-[170px] mt-[10px]">
    </Card.Content>
    <Card.Header class="text-center">
      <Card.Title>{name}</Card.Title>
      <Card.Description>{description}</Card.Description>
    </Card.Header>
    <Card.Footer>
      <p>{address}</p>
    </Card.Footer>
  </Card.Root>
</a>