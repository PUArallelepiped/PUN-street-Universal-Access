<script lang="ts">
	import order_icon from '$lib/assets/order_icon.svg';
	import customer_icon from '$lib/assets/customer_icon.svg';
	import seller_icon from '$lib/assets/seller_icon.svg';
    import { onMount } from 'svelte';
 
    export let firstCol: string;
    export let secondCol: string;
    export let type: string;

    export let ben: boolean = false;
    let self: HTMLDivElement;


    onMount(() => {
        if (ben){
            self.classList.remove("border-PUA-dark-gray");
            self.classList.add("border-PUA-red");
        }
        else{
            self.classList.add("border-PUA-dark-gray");
            self.classList.remove("border-PUA-red");
        }
    });

    function changeButtonStatus() {
        ben = !ben
        if (ben){
            self.classList.remove("border-PUA-dark-gray");
            self.classList.add("border-PUA-red");
        }
        else{
            self.classList.add("border-PUA-dark-gray");
            self.classList.remove("border-PUA-red");
        }
    }
</script>

<div bind:this={self} class="flex my-3 mx-12 justify-between border-2 border-PUA-dark-gray rounded-2xl">
    <button class="w-full hover:bg-gray-200 rounded-xl">
        <div class="flex justify-between items-center">
            <div class="flex items-center">
                {#if type == '0'}
                    <img src={order_icon} alt="" class="h-16 w-16 flex my-4 ml-6" />
                {:else if type == '1'}
                    <img src={customer_icon} alt="" class="h-16 w-16 flex my-4 ml-6" />
                {:else}
                    <img src={seller_icon} alt="" class="h-16 w-16 flex my-4 ml-6" />
                {/if}
                <div class="flex text-left m-4 text-PUA-dark-red">
                    {firstCol}<br>
                    Order user : {secondCol}
                </div>
            </div>
            <div class="m-6">
                {#if type == '0'}
                    <button class="border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red hover:bg-PUA-dark-red hover:text-white">Detail</button>
                {:else}
                    {#if (ben)}
                        <button on:click={changeButtonStatus} class="border-2 border-PUA-dark-red px-7 py-0 text-white bg-PUA-dark-red">Detail</button>
                    {:else}
                        <button on:click={changeButtonStatus} class="border-2 border-PUA-dark-red px-7 py-0 text-PUA-dark-red">Detail</button>
                    {/if}
                {/if}
            </div>
        </div>
    </button>
</div>