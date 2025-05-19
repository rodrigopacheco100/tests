import 'reflect-metadata'
/* eslint-disable no-useless-constructor */
import { container } from 'tsyringe'
import { IRepo, Repository } from './repo'
import { UseCase } from './use-case'

container.register<IRepo>(IRepo.name, Repository)

const useCase1 = container.resolve(UseCase)
const useCase2 = container.resolve(UseCase)

useCase1.run()
useCase2.run()
