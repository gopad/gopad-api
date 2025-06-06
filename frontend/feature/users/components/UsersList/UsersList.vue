<script setup lang="ts">
import {
  FlexRender,
  getCoreRowModel,
  getFilteredRowModel,
  useVueTable,
  type ColumnDef,
  type ColumnFiltersState,
} from '@tanstack/vue-table'
import { useUsers } from '../../providers'
import type { User } from '../../../../client'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import DialogNewUser from '../DialogNewUser'
import { computed, h, ref, unref } from 'vue'
import { valueUpdater } from '@/lib/utils'

const { users, loadUsers } = useUsers()

const columns: ColumnDef<User>[] = [
  {
    accessorKey: 'username',
    header: 'Username',
    cell: ({ row }) => row.getValue('username'),
  },
  {
    accessorKey: 'email',
    header: 'Email',
    cell: ({ row }) => row.getValue('email'),
  },
  {
    accessorKey: 'fullname',
    header: 'Fullname',
    cell: ({ row }) => row.getValue('fullname'),
  },
  {
    accessorKey: 'admin',
    header: 'Admin',
    cell: ({ row }) => (row.getValue('admin') ? 'Yes' : 'No'),
  },
  {
    accessorKey: 'active',
    header: 'Active',
    cell: ({ row }) => (row.getValue('active') ? 'Yes' : 'No'),
  },
  {
    accessorKey: 'created_at',
    header: () => h('div', { class: 'text-end' }, 'Created'),
    cell: ({ row }) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(row.getValue('created_at'))
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  },
  {
    accessorKey: 'updated_at',
    header: () => h('div', { class: 'text-end' }, 'Updated'),
    cell: ({ row }) => {
      const formattedDate = new Intl.DateTimeFormat('en-US').format(
        new Date(row.getValue('updated_at'))
      )

      return h('div', { class: 'text-end' }, formattedDate)
    },
  },
]

const columnFilters = ref<ColumnFiltersState>([])

const table = useVueTable({
  // Using data directly without a getter was not reactive...
  // Might be due to the data coming from context
  // Not really worth the effort to investigate further since the getter works fine
  get data() {
    return unref(users)
  },
  columns,
  getCoreRowModel: getCoreRowModel(),
  getFilteredRowModel: getFilteredRowModel(),
  onColumnFiltersChange: (updaterOrValue) =>
    valueUpdater(updaterOrValue, columnFilters),
  state: {
    get columnFilters() {
      return columnFilters.value
    },
  },
})

const filter = computed({
  get() {
    return table.getColumn('name')?.getFilterValue() as string
  },
  set(value: string) {
    table.getColumn('name')?.setFilterValue(value)
  },
})

loadUsers()
</script>

<template>
  <div class="w-full">
    <div class="flex gap-2 items-center justify-between py-4">
      <Input v-model="filter" class="max-w-sm" placeholder="Filter" />
      <DialogNewUser />
    </div>
    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow
            v-for="headerGroup in table.getHeaderGroups()"
            :key="headerGroup.id"
          >
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender
                v-if="!header.isPlaceholder"
                :render="header.column.columnDef.header"
                :props="header.getContext()"
              />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <template v-for="row in table.getRowModel().rows" :key="row.id">
              <TableRow :data-state="row.getIsSelected() && 'selected'">
                <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                  <FlexRender
                    :render="cell.column.columnDef.cell"
                    :props="cell.getContext()"
                  />
                </TableCell>
              </TableRow>
              <TableRow v-if="row.getIsExpanded()">
                <TableCell :colspan="row.getAllCells().length">
                  {{ JSON.stringify(row.original) }}
                </TableCell>
              </TableRow>
            </template>
          </template>

          <TableRow v-else>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              No results
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  </div>
</template>
