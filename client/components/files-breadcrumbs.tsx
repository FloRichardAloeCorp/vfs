import { BreadcrumbItem, Breadcrumbs } from '@nextui-org/breadcrumbs';
import * as React from 'react';

export interface IFilesBreadCrumbsProps {
    onItemClicked: (item: string) => void
    path: string
    className?: string
}

export function FilesBreadCrumbs(props: IFilesBreadCrumbsProps) {
    const cutPathFrom = (part: string) => {
        let newPath = '/'
        if (part !== '') {
            const pathParts = props.path.split('/')
            const partIndex = pathParts.indexOf(part)
            newPath = pathParts.slice(0, partIndex + 1).join('/')
        }

        props.onItemClicked(newPath)
    }

    const breadcrumbsItems = props.path.split('/').map(part => {
        return <BreadcrumbItem key={crypto.randomUUID()} onClick={() => cutPathFrom(part)} >{part}</BreadcrumbItem>
    })

    return (
        <Breadcrumbs separator="/" itemClasses={{ separator: "px-2" }} className={props.className}>
            {breadcrumbsItems}
        </Breadcrumbs>
    );
}
