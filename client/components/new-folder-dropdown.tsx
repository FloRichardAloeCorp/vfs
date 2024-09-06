import { Button } from '@nextui-org/button';
import { Dropdown, DropdownItem, DropdownMenu, DropdownSection, DropdownTrigger } from '@nextui-org/dropdown';
import { Input } from '@nextui-org/input';
import path from 'path';
import * as React from 'react';
import { useRef, useState } from 'react';

export interface INewFolderDropdownProps {
    isOpen?: boolean
    onNewFolder: (newPath: string) => void
    currentPath: string
}

export function NewFolderDropdown(props: INewFolderDropdownProps) {
    const [dropdownOpen, setDropdownOpen] = useState(props.isOpen)
    const folderNameInputRef = useRef<HTMLInputElement>(null)

    const newFolder = async (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key != "Enter") {
            return
        }

        const folderPath = path.join(props.currentPath, folderNameInputRef.current?.value as string)
        try {
            await fetch(`${process.env.NEXT_PUBLIC_VFS_BASE_URL}/directory${folderPath}`, { method: "POST" })
        } catch (error) {
            console.log(error)
        } finally {
            setDropdownOpen(false)
            props.onNewFolder(folderPath)
        }
    }


    return (
        <Dropdown isOpen={dropdownOpen} closeOnSelect={false} onClose={() => setDropdownOpen(false)}>
            <DropdownTrigger>
                <Button variant='light' className='min-h-0 min-w-0 h-fit px-0' onClick={() => setDropdownOpen(true)}>
                    New folder
                </Button>
            </DropdownTrigger>
            <DropdownMenu variant='light'>
                <DropdownSection title="Folder name">
                    <DropdownItem onClick={() => { folderNameInputRef.current?.focus() }} description="Press enter to validate">
                        <Input type='text' ref={folderNameInputRef} className='pb-2' onKeyDown={newFolder}></Input>
                    </DropdownItem>
                </DropdownSection>
            </DropdownMenu>
        </Dropdown >
    );
}
