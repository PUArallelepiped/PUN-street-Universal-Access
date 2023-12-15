<script lang="ts">
	import admin_icon from '$lib/assets/admin_icon.svg';
    import AdminOrderCart from '$lib/components/PUA/adminOrder.svelte'

    let orders: {
        date: string,
        user: string
    }[] = [
        {
            date: "2020-10-11 19:20:50",
            user: "user01"
        },
        {
            date: "2020-10-11 19:20:50",
            user: "user02"
        },
        {
            date: "2020-10-11 19:20:50",
            user: "user03"
        },
        {
            date: "2020-10-11 19:20:50",
            user: "user04"
        }
    ]

    let users: {
        name: string,
        email: string
    }[] = [
        {
            name: "sakana",
            email: "r1111111"
        },
        {
            name: "sakana dewanai",
            email: "t1111111"
        },
        {
            name: "sakana desu",
            email: "fsddsdss"
        },
    ]
</script>

<div class="flex justify-center font-bold">
    <div class="flex w-3/5 flex-col">
        <div class="flex bg-PUA-stone">
            <img src={admin_icon} alt="" class="h-28 w-28 flex my-6 ml-10" />
            <div class="flex flex-col m-7 justify-between">
                <div>
                    <div class="text-PUA-gray text-2xl">sakana san</div>
                    <div class="text-gray-300">t111105911</div>
                </div>
                <div class="text-PUA-gray text-xl">address</div>
            </div>
        </div>
        <div class="flex h-10">
            <button class="w-full bg-slate-400">aa</button>
            <button class="w-full bg-slate-100">aa</button>
        </div>
        <div class="h-screen bg-white">
            {#each orders as order}
                <AdminOrderCart firstCol= {order.date} secondCol = {order.user}></AdminOrderCart>
            {/each}
        </div>
        <div class="h-screen bg-white">
            {#each users as user}
                <AdminOrderCart firstCol= {user.name} secondCol = {user.email}></AdminOrderCart>
            {/each}
        </div>
    </div>
</div>